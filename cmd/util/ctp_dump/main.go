package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/mineralres/goshare/pkg/util"
)

// CtpPrimitiveType mmber
type CtpPrimitiveType struct {
	typeName         string
	cppPrimitiveType string // char, char array , int , short, double
	arrayLen         uint32
	cppSource        string
}

// CtpTypeMember member
type CtpTypeMember struct {
	cppSource  string
	memberName string
	primitive  CtpPrimitiveType
}

// CtpCompoundType  ctp复合类型
type CtpCompoundType struct {
	cppTypeName string          // 原始名称, 如TThostFtdcExchangePropertyType
	cppSource   string          // 原始代码
	members     []CtpTypeMember // 成员
}

// CtpClassMethodType ctp中class的方法
type CtpClassMethodType struct {
	cppSource  string // 如 virtual void OnFrontConnected(){};
	methodName string // 如 Join, OnRspSubForQuoteRsp, OnRtnDepthMarketData, RegisterFront 等
}

// CtpClassType ctp中的几个class
type CtpClassType struct {
	cppTypeName string                // 原始名称, 如CThostFtdcMdSpi
	cppSource   string                // 原始代码, 如class CThostFtdcMdSpi
	methods     []*CtpClassMethodType // 方法列表
}

// 一些define定义的类型
func parseDataType(header string) ([]CtpPrimitiveType, error) {
	reg := regexp.MustCompile(`(\[[0-9]{1,5}\])`)
	reg2 := regexp.MustCompile(`\[([0-9]{1,5})\]`)
	var typelist []CtpPrimitiveType
	file, err := os.Open(header)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	start := false
	for scanner.Scan() {
		line := util.StringFromGBK(scanner.Text())
		line = strings.Trim(line, "\t")
		if len(line) == 0 {
			continue
		}
		var ct = CtpPrimitiveType{cppSource: line}
		if !start {
			if strings.Index(ct.cppSource, "typedef") == 0 {
				start = true
			} else {
				continue
			}
		}
		if strings.Index(ct.cppSource, "using") == 0 {
			continue
		}
		if strings.Index(ct.cppSource, "namespace") == 0 {
			continue
		}
		if strings.Index(ct.cppSource, "///") == 0 {
			continue
		}
		if strings.Index(ct.cppSource, "#endif") == 0 {
			continue
		}
		if strings.Index(ct.cppSource, "#define") == 0 {
			// #define THOST_FTDC_EXP_Normal '0'
			items := strings.Split(ct.cppSource, " ")
			ct.typeName = items[1]
		}
		if strings.Index(ct.cppSource, "typedef") == 0 {
			// typedef char TThostFtdcSystemIDType[21];
			items := strings.Split(ct.cppSource, " ")
			switch items[1] {
			case "char", "double", "int", "short":
			default:
				panic("未处理的类型")
			}
			ct.cppPrimitiveType = items[1]
			ct.typeName = strings.ReplaceAll(items[2], ";", "")
			l := reg2.FindStringSubmatch(ct.typeName)
			if len(l) > 0 {
				ct.arrayLen = uint32(util.ParseInt(l[1]))
			}
			ct.typeName = reg.ReplaceAllString(ct.typeName, "")
			if strings.Index(ct.typeName, "[") > 0 {
				log.Println("ct", ct.cppPrimitiveType, ct.typeName)
				panic("处理异常")
			}
		}
		ct.typeName = strings.TrimSpace(ct.typeName)
		if ct.cppSource == "}" {
			// 不push
			continue
		}
		typelist = append(typelist, ct)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return typelist, nil
}

// 复合类型
func parseAPIStruct(header string, primitiveTypeList []CtpPrimitiveType) ([]CtpCompoundType, error) {
	var typelist []CtpCompoundType
	file, err := os.Open(header)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	start := false
	for scanner.Scan() {
		line := util.StringFromGBK(scanner.Text())
		line = strings.Trim(line, "\t")
		if len(line) == 0 {
			continue
		}
		if !start {
			if strings.Index(line, "struct") == 0 {
				start = true
			} else {
				continue
			}
		}
		if strings.Index(line, "using") == 0 {
			continue
		}
		if strings.Index(line, "///") == 0 {
			continue
		}
		if strings.Index(line, "#endif") == 0 {
			continue
		}
		if strings.Index(line, "struct") == 0 {
			var ct = CtpCompoundType{cppSource: line}
			ct.cppTypeName = strings.Replace(line, "struct", "", -1)
			ct.cppTypeName = strings.Replace(ct.cppTypeName, "{", "", -1)
			ct.cppTypeName = strings.Replace(ct.cppTypeName, " ", "", -1)
			typelist = append(typelist, ct)
		} else {
			l := len(typelist)
			if line != "}" {
				typelist[l-1].cppSource += line
			}
			if line == "{" || line == "};" || line == "}" {

			} else {
				f := false
				for _, p := range primitiveTypeList {
					if strings.Index(line, p.typeName) >= 0 {
						items := strings.Split(line, "\t")
						items[1] = strings.ReplaceAll(items[1], ";", "")
						// log.Println(line, items)
						typelist[l-1].members = append(typelist[l-1].members, CtpTypeMember{primitive: p, cppSource: line, memberName: items[1]})
						f = true
						break
					}
				}
				if !f {
					log.Println("", line, len(primitiveTypeList))
					panic("struct .member not found in primitiveTypeList")
				}
			}
		}

	}
	return typelist, nil
}

// API和SPI class
func parseClass(header string) []*CtpClassType {
	reg := regexp.MustCompile(`[\s\*]\w+[(]`)
	var classlist []*CtpClassType
	file, err := os.Open(header)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var ct *CtpClassType
	for scanner.Scan() {
		line := util.StringFromGBK(scanner.Text())
		line = strings.Trim(line, "\t")
		if len(line) == 0 {
			continue
		}
		if strings.Index(line, "class") == 0 {
			// 类型
			ct = &CtpClassType{cppSource: line}
			ct.cppSource = line
			arr := strings.Split(line, " ")
			ct.cppTypeName = arr[len(arr)-1]
			classlist = append(classlist, ct)
		} else if strings.Index(line, "virtual") == 0 || strings.Index(line, "static") == 0 {
			// 方法
			method := &CtpClassMethodType{cppSource: line}
			arr := reg.FindStringSubmatch(line)
			if len(arr) > 0 {
				method.methodName = arr[0]
				method.methodName = strings.TrimFunc(method.methodName, func(r rune) bool {
					if r == '(' || r == '*' || r == ' ' {
						return true
					}
					return false
				})
				log.Println(method.methodName)
			} else {
				log.Println(line)
				panic("")
			}
			ct.methods = append(ct.methods, method)
		} else {
			continue
		}
	}
	return classlist
}

func writeGoFile(primitiveTypeList []CtpPrimitiveType, typList []CtpCompoundType, path, pkgName string) error {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf(`package %s

import (
	"encoding/binary"
	"io"
	"strings"
)

func writeLen(w io.Writer, source string, fixedLen int) {
	if len(source) < fixedLen {
		io.WriteString(w, source)
		b := make([]byte, fixedLen-len(source))
		w.Write(b)
		return
	}
	io.WriteString(w, source[:fixedLen])
}

func readLen(r io.Reader, fixedLen int) string {
	buf := make([]byte, fixedLen)
	n, err := r.Read(buf)
	if err != nil || n != fixedLen {
		return ""
	}
	return strings.Trim(string(buf), "\u0000")
}

`, pkgName))
	for _, t := range typList {
		members := ""
		marshal := ""
		unmarshal := ""
		for _, m := range t.members {
			members += "    " + m.memberName + " "
			switch m.primitive.cppPrimitiveType {
			case "char":
				if m.primitive.arrayLen > 0 {
					members += "string"
					marshal += fmt.Sprintf("    writeLen(w, m.%s, %d)\r\n", m.memberName, m.primitive.arrayLen)
					unmarshal += fmt.Sprintf("    m.%s = readLen(r, %d)\r\n", m.memberName, m.primitive.arrayLen)
				} else {
					members += "uint8"
					marshal += fmt.Sprintf("    binary.Write(w, binary.BigEndian, &m.%s)\r\n", m.memberName)
					unmarshal += fmt.Sprintf("    binary.Read(r, binary.BigEndian, &m.%s)\r\n", m.memberName)
				}
			case "int":
				members += "int32"
				marshal += fmt.Sprintf("    binary.Write(w, binary.BigEndian, &m.%s)\r\n", m.memberName)
				unmarshal += fmt.Sprintf("    binary.Read(r, binary.BigEndian, &m.%s)\r\n", m.memberName)
			case "short":
				members += "int16"
				marshal += fmt.Sprintf("    binary.Write(w, binary.BigEndian, &m.%s)\r\n", m.memberName)
				unmarshal += fmt.Sprintf("    binary.Read(r, binary.BigEndian, &m.%s)\r\n", m.memberName)
			case "double":
				members += "float64"
				marshal += fmt.Sprintf("    binary.Write(w, binary.BigEndian, &m.%s)\r\n", m.memberName)
				unmarshal += fmt.Sprintf("    binary.Read(r, binary.BigEndian, &m.%s)\r\n", m.memberName)
			}
			members += ";\r\n"
		}
		f.WriteString(fmt.Sprintf("type %s struct {\r\n%s}\r\n\r\n", t.cppTypeName, members))
		// marshal + unmarshal
		f.WriteString(fmt.Sprintf("func (m*%s) Marshal(w io.Writer) {\r\n%s}\r\n\r\n", t.cppTypeName, marshal))
		f.WriteString(fmt.Sprintf("func (m*%s) Unmarshal(r io.Reader) {\r\n%s}\r\n\r\n", t.cppTypeName, unmarshal))
	}
	return err
}

func writeJSBinaryParser(primitiveTypeList []CtpPrimitiveType, typList []CtpCompoundType, path, pkgName string) error {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(`var Parser = require("binary-parser-encoder").Parser;
const stropt = n => {
	return {
		stripNull: true,
		length: n
	}
}

	`)
	for _, t := range typList {
		src := ""
		initSrc := ""
		sz := 0
		max := 1
		for _, m := range t.members {
			switch m.primitive.cppPrimitiveType {
			case "char":
				initSrc += fmt.Sprintf("%s:'', ", m.memberName)
				if m.primitive.arrayLen > 0 {
					if m.memberName == "InstrumentName" || m.memberName == "StatusMsg" {
						src += fmt.Sprintf(".array('%s', {type:'uint8', length: %d})", m.memberName, m.primitive.arrayLen)
					} else {
						src += fmt.Sprintf(".string('%s', stropt(%d))", m.memberName, m.primitive.arrayLen)
					}
					sz += int(m.primitive.arrayLen)
				} else {
					src += fmt.Sprintf(".string('%s', stropt(1))", m.memberName)
					sz++
				}
			case "int":
				initSrc += fmt.Sprintf("%s:0, ", m.memberName)
				skip := sz % 4
				if skip > 0 {
					skip = 4 - skip
					src += fmt.Sprintf(".skip(%d).int32le('%s')", skip, m.memberName)
				} else {
					src += fmt.Sprintf(".int32le('%s')", m.memberName)
				}
				sz += (4 + skip)
				if max < 4 {
					max = 4
				}
				// log.Println(m.memberName, skip, sz)
			case "short":
				initSrc += fmt.Sprintf("%s:0, ", m.memberName)
				skip := (2 - sz%2)
				if skip > 0 {
					skip = 2 - skip
					src += fmt.Sprintf(".skip(%d).int16le('%s')", skip, m.memberName)
				} else {
					src += fmt.Sprintf(".int16le('%s')", m.memberName)
				}
				sz += (2 + skip)
				if max < 2 {
					max = 2
				}
			case "double":
				initSrc += fmt.Sprintf("%s:0, ", m.memberName)
				skip := sz % 8
				if skip > 0 {
					skip = 8 - skip
					src += fmt.Sprintf(".skip(%d).doublele('%s')", skip, m.memberName)
				} else {
					src += fmt.Sprintf(".doublele('%s')", m.memberName)
				}
				sz += (8 + skip)
				if max < 8 {
					max = 8
				}
				// log.Println(m.memberName, skip, sz)
			}
		}
		skip := sz % max
		if skip > 0 {
			f.WriteString(fmt.Sprintf("exports.%sParser = () => new Parser()%s.skip(%d);\r\n", t.cppTypeName, src, max-skip))
			sz += max - skip
		} else {
			f.WriteString(fmt.Sprintf("exports.%sParser = () => new Parser()%s;\r\n", t.cppTypeName, src))
		}
		f.WriteString(fmt.Sprintf("exports.%sInit = () => {return {%s}};\r\n", t.cppTypeName, initSrc))
		if t.cppTypeName == "CThostFtdcDepthMarketDataField" {
			log.Println("SIZE = ", sz)
		}
	}
	return nil
}

func writeAddonHelper(tp, header, output string) error {
	classList := parseClass(header)
	f, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	c := classList[0]
	f.WriteString("\r\n\r\n\r\n // CTP的回调函数 再回调js---------------------- \r\n")
	for _, m := range c.methods {
		callSrc := ""
		if strings.Index(m.cppSource, "OnFrontConnected") > 0 {
			callSrc = `callJS(new Message("OnFrontConnected", "", 0));`
		} else if strings.Index(m.cppSource, "OnFrontDisconnected") > 0 {
			callSrc = `callJS(new Message("OnFrontDisconnected", "", 0));`
		} else if strings.Index(m.cppSource, "OnHeartBeatWarning") > 0 {
			callSrc = fmt.Sprintf(`callJS(new Message("%s", "", 0));`, m.methodName)
		} else if strings.Index(m.cppSource, "void OnRsp") > 0 {
			reg := regexp.MustCompile(`[\*]\w+[,]`)
			arr := reg.FindStringSubmatch(m.cppSource)
			arr[0] = arr[0][1:]
			t1 := arr[0]
			callSrc = fmt.Sprintf(`    callJS(make_message("%s", %s pRspInfo, nRequestID, bIsLast)); `, m.methodName, t1)
		} else if strings.Index(m.cppSource, "void OnRtn") > 0 {
			// 一个参数的，如OnRtnOrder
			reg := regexp.MustCompile(`[\*]\w+[\)]`)
			arr := reg.FindStringSubmatch(m.cppSource)
			arr[0] = arr[0][1:]
			t1 := arr[0]
			t1 = strings.ReplaceAll(t1, ")", "")
			callSrc = fmt.Sprintf(`    callJS(make_message("%s", %s)); `, m.methodName, t1)
		} else if strings.Index(m.cppSource, "void OnErrRtn") > 0 {
			// 两个参数的,如OnErrRtnOrderInsert
			reg := regexp.MustCompile(`[\*]\w+[,]`)
			arr := reg.FindStringSubmatch(m.cppSource)
			arr[0] = arr[0][1:]
			t1 := arr[0]
			t1 = strings.ReplaceAll(t1, ",", "")
			callSrc = fmt.Sprintf(`    callJS(make_message("%s", %s, pRspInfo)); `, m.methodName, t1)
		}
		src := strings.Replace(m.cppSource, "{};", fmt.Sprintf(`{ %s }
		`, callSrc), -1)
		src = strings.ReplaceAll(src, "virtual ", "")
		f.WriteString(src)
	}

	// 以下是生成js call Api的方法
	c = classList[1]
	for _, m := range c.methods {
		if m.methodName == "CreateFtdcTraderApi" || m.methodName == "CreateFtdcMdApi" {
			continue
		}
		src := fmt.Sprintf("InstanceMethod(\"%s\", &%s::%s),\r\n", m.methodName, tp, m.methodName)
		f.WriteString(src)
	}
	f.WriteString("\r\n\r\n\r\n // 以下是生成js call Api的方法---------------------- \r\n")
	for _, m := range c.methods {
		if m.methodName == "RegisterSpi" || m.methodName == "CreateFtdcTraderApi" || m.methodName == "CreateFtdcMdApi" {
			continue
		}
		src := ""
		if m.methodName == "Release" || m.methodName == "Init" || m.methodName == "Join" {
			src = fmt.Sprintf(`
			void %s(const CallbackInfo &info) { api_->%s(); }
			`, m.methodName, m.methodName)
		} else if m.methodName == "RegisterFront" || m.methodName == "RegisterNameServer" {
			src = fmt.Sprintf(`
			void %s(const CallbackInfo &info) {
				std::string str = info[0].As<Napi::String>();
				api_->%s((char*)str.data());
			}
			`, m.methodName, m.methodName)
		} else if m.methodName == "SubscribePrivateTopic" || m.methodName == "SubscribePublicTopic" {
			src = fmt.Sprintf(`
			void %s(const CallbackInfo &info) {
				int32_t nResumeType = info[0].As<Napi::Number>().Int32Value();
				api_->%s((THOST_TE_RESUME_TYPE)nResumeType);
			}
			`, m.methodName, m.methodName)
		} else if m.methodName == "GetApiVersion" || m.methodName == "GetTradingDay" {
			src = fmt.Sprintf(`
			Napi::Value %s(const CallbackInfo &info) {return Napi::String::New(info.Env(), api_->%s()); }
			`, m.methodName, m.methodName)
		} else if strings.Index(m.cppSource, "char *ppInstrumentID[], int nCount") > 0 {
			src = fmt.Sprintf(`
			Napi::Value %s(const CallbackInfo &info) {
				auto arr = info[0].As<Napi::Array>();
				auto instruments = new char *[arr.Length()];
				for (size_t i = 0; i < arr.Length(); i++) {
					Napi::Value v = arr[i];
					std::string str = v.As<Napi::String>();
					instruments[i] = new char[str.size()];
					memcpy(instruments[i], str.c_str(), str.size());
				}
				int ret = api_->%s(instruments, arr.Length());
				for (size_t i = 0; i < arr.Length(); i++) {
					delete instruments[i];
				}
				delete instruments;
				return Napi::Number::New(info.Env(), ret);
			}
			`, m.methodName, m.methodName)
			log.Println("订阅", m.methodName, src)
		} else {
			reg := regexp.MustCompile(`[(]\w+[\s\*]`)
			arr := reg.FindStringSubmatch(m.cppSource)
			arr[0] = arr[0][1:]
			t1 := arr[0]

			if strings.Index(m.cppSource, "nRequestID") > 0 {
				// 两个参数的
				src = fmt.Sprintf(`
			void %s(const CallbackInfo &info) {
				auto ab = info[0].As<Napi::TypedArray>().ArrayBuffer();
				int32_t nRequestID = info[1].As<Napi::Number>().Int32Value();
				%s* req = (%s*)ab.Data();
				api_->%s(req, nRequestID);
			}
			`, m.methodName, arr[0], arr[0], m.methodName)
			} else {
				// 一个参数的
				log.Println(m.methodName, m.cppSource)
				src = fmt.Sprintf(`
			void %s(const CallbackInfo &info) {
				auto ab = info[0].As<Napi::TypedArray>().ArrayBuffer();
				%s* req = (%s*)ab.Data();
				api_->%s(req);
			}
			`, m.methodName, t1, t1, m.methodName)
			}
		}

		f.WriteString(src)
	}
	return nil
}

// CTP future api headers
func parseCtpFutureAPIHeaders() error {
	l, err := parseDataType("ctp_future_headers//ThostFtdcUserApiDataType.h")
	l2, err := parseAPIStruct("ctp_future_headers//ThostFtdcUserApiStruct.h", l)
	for _, i := range l2 {
		log.Println(i)
	}
	// writeGoFile(l, l2, "output/ctp_business.go", "ftdc")
	writeJSBinaryParser(l, l2, "output/ctp_business.txt", "ctp")
	return err
}

// ctp sopt api headers
func parseCtpSoptAPIHeaders() error {
	l, err := parseDataType("ctp_sopt_headers//ThostFtdcUserApiDataType.h")
	l2, err := parseAPIStruct("ctp_sopt_headers//ThostFtdcUserApiStruct.h", l)
	writeGoFile(l, l2, "output/ctp_business.go", "ftdc")
	return err
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	parseCtpFutureAPIHeaders()
	return
	// parseCtpSoptAPIHeaders()
	writeAddonHelper("MdApi", "ctp_future_headers//ThostFtdcMdApi.h", "output/md-addon.txt")
	// writeAddonHelper("TraderApi", "ctp_future_headers//ThostFtdcTraderApi.h", "output/trade-addon.txt")
	// for _, v := range l {
	// 	log.Println(v)
	// 	for _, m := range v.methods {
	// 		log.Println(m)
	// 	}
	// }
}

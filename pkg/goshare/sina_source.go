package goshare

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mineralres/goshare/pkg/base"
	"github.com/mineralres/goshare/pkg/pb"
)

// GetLastTick 取最新行情
func (p *SinaSource) GetLastTick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	if symbol.Exchange == pb.ExchangeType_SSE || symbol.Exchange == pb.ExchangeType_SZE {
		if symbol.Exchange == pb.ExchangeType_SSE && strings.Index(symbol.Code, "1000") == 0 {
			// 上证50ETF期权tick
			return getSSEOptionTick(symbol)
		}
		// 股票tick
		return getStockLastTick(symbol)
	}

	if symbol.Exchange == pb.ExchangeType_INDEX {
		// 指数tick
		return p.GetIndexLastTick(symbol)
	}
	return nil, base.ErrUnsupported
}

// GetKData 请求历史K线数据
/*
symbol：股票代码，即6位数字代码，或者指数代码
startTime：开始时间time_t
endTime：结束时间time_t
period：周期
retryCount：当网络异常后重试次数，默认为3
*/
func (p *SinaSource) GetKData(symbol *pb.Symbol, period pb.PeriodType, startTime, endTime int64, retryCount int) (*pb.KlineSeries, error) {
	ex := symbol.Exchange
	if ex == pb.ExchangeType_SSE || ex == pb.ExchangeType_SZE {
		var svc EaseMoneySource
		// 股票K线
		return svc.GetCNStockKData(symbol, period, startTime, endTime, retryCount)

	} else if ex == pb.ExchangeType_SHFE || ex == pb.ExchangeType_CZCE || ex == pb.ExchangeType_DCE || ex == pb.ExchangeType_CFFEX || ex == pb.ExchangeType_INE {
		// 期货K线
		return getCNFutureKData(symbol, period, startTime, endTime, retryCount)

	} else if ex == pb.ExchangeType_OPTION_SSE {
		// 期权K线
		return getOptionSSEKData(symbol, period, startTime, endTime, retryCount)

	}
	var ret pb.KlineSeries
	return &ret, base.ErrUnsupported
}

// getStockLastTick 取股票最新报价
func getStockLastTick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	ret := &pb.MarketDataSnapshot{}
	exstr := "sh"
	if symbol.Exchange == pb.ExchangeType_SZE {
		exstr = "sz"
	}
	tickArr := getRawTickString(exstr, symbol.Code)
	if tickArr == nil || len(tickArr) < 38 {
		return nil, errors.New("ErrGetStockTick")
	}
	if tickArr != nil && len(tickArr) >= 38 {
		timeStr := tickArr[30]
		tx, err := time.ParseInLocation("20060102150405", timeStr, time.Local)
		if err == nil {
			ret.Time = tx.Unix()
		}
		td, err := strconv.Atoi(time.Unix(ret.Time, 0).Format("20060102"))
		if err == nil {
			ret.TradingDay = int32(td)
		}

		ret.Symbol = *symbol
		ret.Price = base.ParseFloat(tickArr[3])
		ret.Close = ret.Price
		ret.PreClose = base.ParseFloat(tickArr[4])
		ret.Open = base.ParseFloat(tickArr[5])
		ret.High = base.ParseFloat(tickArr[33])
		ret.Low = base.ParseFloat(tickArr[34])
		ret.Volume = (base.ParseFloat(tickArr[6]))
		ret.Amount = float64(base.ParseInt(tickArr[37]) * 10000)
		ret.UpperLimitPrice = base.ParseFloat(tickArr[47])
		ret.LowerLimitPrice = base.ParseFloat(tickArr[48])
		var ob5 pb.OrderBook
		ob5.BidVolume = base.ParseFloat(tickArr[18])
		ob5.Bid = base.ParseFloat(tickArr[17])
		ob5.AskVolume = base.ParseFloat(tickArr[28])
		ob5.Ask = base.ParseFloat(tickArr[27])
		var ob4 pb.OrderBook
		ob4.BidVolume = base.ParseFloat(tickArr[16])
		ob4.Bid = base.ParseFloat(tickArr[15])
		ob4.AskVolume = base.ParseFloat(tickArr[26])
		ob4.Ask = base.ParseFloat(tickArr[25])
		var ob3 pb.OrderBook
		ob3.BidVolume = base.ParseFloat(tickArr[14])
		ob3.Bid = base.ParseFloat(tickArr[13])
		ob3.AskVolume = base.ParseFloat(tickArr[24])
		ob3.Ask = base.ParseFloat(tickArr[23])
		var ob2 pb.OrderBook
		ob2.BidVolume = base.ParseFloat(tickArr[12])
		ob2.Bid = base.ParseFloat(tickArr[11])
		ob2.AskVolume = base.ParseFloat(tickArr[22])
		ob2.Ask = base.ParseFloat(tickArr[21])
		var ob1 pb.OrderBook
		ob1.BidVolume = base.ParseFloat(tickArr[10])
		ob1.Bid = base.ParseFloat(tickArr[9])
		ob1.AskVolume = base.ParseFloat(tickArr[20])
		ob1.Ask = base.ParseFloat(tickArr[19])
		ret.OrderBookList = []pb.OrderBook{ob1, ob2, ob3, ob4, ob5}
	}
	return ret, nil
}

func getRawTickString(exstr string, symbol string) []string {
	resp, err := http.Get("http://web.sqt.gtimg.cn/q=" + exstr + symbol)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			tickArr := strings.Split(string(body), "~")
			if len(tickArr) > 1 {
				tickArr[1] = base.StringFromGBK(tickArr[1])
			} else {
				log.Printf("getRawTickString %s-%s", exstr, symbol)
			}
			return tickArr
		}
	}
	return nil
}

// GetIndexLastTick 指数行情
func (p *SinaSource) GetIndexLastTick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	ret := &pb.MarketDataSnapshot{}
	resp, err := http.Get("http://hq.sinajs.cn/list=" + symbol.Code)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			tickArr := strings.Split(string(body), ",")
			//sym := strings.Split(string(tickArr[0]), "=")
			ret.Symbol = *symbol
			ret.Price = base.ParseFloat(tickArr[1])
			ret.Close = ret.Price
			return ret, nil
		}
	}
	return nil, errors.New("ErrGetIndex")
}

// GetMainFutureLastTick 取主力合约
func (p *SinaSource) GetMainFutureLastTick(et pb.ExchangeType) ([]pb.MarketDataSnapshot, error) {
	var ret []pb.MarketDataSnapshot
	var etStr string
	switch et {
	case pb.ExchangeType_SHFE:
		etStr = "SHFE"
	case pb.ExchangeType_CZCE:
		etStr = "CZCE"
	case pb.ExchangeType_DCE:
		etStr = "DCE"
	case pb.ExchangeType_CFFEX:
		etStr = "_168"
	default:
		return ret, fmt.Errorf("error ExchangeType %s", et)
	}

	address := fmt.Sprintf("http://nufm.dfcfw.com/EM_Finance2014NumericApplication/JS.aspx?type=CT&cmd=C.%s", etStr) + "&sty=FCFL4O&sortType=(ChangePercent)&sortRule=-1&page=1&pageSize=200&js={rank:[(x)],pages:(pc),total:(tot)}&token=7bc05d0d4c3c22ef9fca8c2a912d779c&jsName=quote_123&_g=0.628606915911589&_=1521620666159"

	resp, err := http.Get(address)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err == nil {
			tickArr := strings.Split(string(body), "\"")
			// fmt.Println(string(body))
			i := 0
			for i < len(tickArr) {

				mktStrArr := strings.Split(string(tickArr[i]), ",")
				i++
				if len(mktStrArr) < 15 {
					continue
				}
				if len(mktStrArr[1]) > 3 {
					continue
				}

				mkt := pb.MarketDataSnapshot{}
				mkt.Symbol = pb.Symbol{Exchange: et, Code: mktStrArr[1]}
				mkt.Open = base.ParseFloat(mktStrArr[11])
				mkt.High = base.ParseFloat(mktStrArr[13])
				mkt.Low = base.ParseFloat(mktStrArr[14])
				mkt.Price = base.ParseFloat(mktStrArr[3])
				mkt.Close = mkt.Price
				mkt.Volume = base.ParseFloat(mktStrArr[10])
				mkt.Amount = base.ParseFloat(mktStrArr[15])
				mkt.Position = base.ParseFloat(mktStrArr[9])
				mkt.PreSettlementPrice = base.ParseFloat(mktStrArr[8])
				ret = append(ret, mkt)
				// fmt.Println(mkt.Symbol, mkt.Open, mkt.High, mkt.Low)
			}
		}

	}
	return ret, nil
}

// GetIndexMember 指数成份股
func (p *SinaSource) GetIndexMember(symbol *pb.Symbol, retryCount int) ([]pb.Symbol, error) {
	return getIndexMem(symbol)
}

func getIndexMem(symbol *pb.Symbol) ([]pb.Symbol, error) {
	var ret []pb.Symbol

	page_number := 1
	member_number := 0

	for true {
		address := fmt.Sprintf("http://vip.stock.finance.sina.com.cn/corp/view/vII_NewestComponent.php?page=%d&indexid=%s", page_number, symbol.Code)
		// log.Println(address)
		page_number++
		doc, err := goquery.NewDocument(address)

		if err != nil {
			fmt.Println(err)
			return ret, err
		}

		b_empty := true

		doc.Find("#NewStockTable").Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
			if i > 0 {
				b_empty = false
				code := s.Find("div").Eq(0).Text()
				// fmt.Println(code)
				s, err := formatSymbol(code)
				if err == nil {
					ret = append(ret, s)
					member_number += 1
				}
			}
		})
		if b_empty == true || doc.Find("#page_form").Length() == 0 {
			break
		}
	}

	// log.Println(ret)
	return ret, nil
}

func formatSymbol(code string) (pb.Symbol, error) {
	var ret pb.Symbol
	if len(code) < 6 {
		return ret, fmt.Errorf("error code %s", code)
	}

	switch code[0] {
	case '6':
		return pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: code}, nil
	case '0':
		return pb.Symbol{Exchange: pb.ExchangeType_SZE, Code: code}, nil
	case '3':
		return pb.Symbol{Exchange: pb.ExchangeType_SZE, Code: code}, nil
	default:
		return ret, fmt.Errorf("error code %s", code)
	}
}

func adaptCZCE(value string) string {
	if strings.Index(value, "D-") >= 0 {
		value = value[2:]
	}
	for i := 0; i < len(value); i++ {
		if value[i] >= '0' && value[i] <= '9' && i > 0 {
			if len(value)-i == 3 {
				return value[0:i] + "1" + value[i:]
			}
			break
		}
	}
	return value
}

func getCNFutureKData(symbol *pb.Symbol, period pb.PeriodType, startTime, endTime int64, retryCount int) (*pb.KlineSeries, error) {
	var ret pb.KlineSeries
	type SinaKline struct {
		ClosePrice string `json:"c"`
		Day        string `json:"d"`
		MaxPrice   string `json:"h"`
		MinPrice   string `json:"l"`
		NowVolume  string `json:"v"`
		OpenPrice  string `json:"o"`
	}

	isDaily := false
	ktype := "5"
	switch period {
	case pb.PeriodType_D1:
		isDaily = true
		ktype = ""
	case pb.PeriodType_M1:
		ktype = "1"
	case pb.PeriodType_M5:
		ktype = "5"
	case pb.PeriodType_H1:
		ktype = "60"
	}
	code := symbol.Code
	qapi := "http://stock2.finance.sina.com.cn/futures/api/jsonp.php//InnerFuturesNewService.getFewMinLine?symbol=" + adaptCZCE(code) + "&type=" + ktype
	qapi = fmt.Sprintf("https://stock.sina.com.cn/futures/api/jsonp.php/var_X=/InnerFuturesNewService.getFewMinLine?symbol=%s&type=%s", adaptCZCE(code), ktype)
	if ktype == "" {
		tradingDay := time.Now().Format("20060102")
		qapi = fmt.Sprintf("https://stock.sina.com.cn/futures/api/jsonp.php/var_X=/InnerFuturesNewService.getDailyKLine?symbol=%s&_=%s", adaptCZCE(code), tradingDay)
		isDaily = true
	}
	resp, err := http.Get(qapi)
	if err != nil {
		return &ret, err
	}
	v, err := ioutil.ReadAll(resp.Body)
	str := strings.TrimLeft(string(v), "var_X=(")
	xl := len(str)
	if xl > 2 && err == nil {
		dataStr := string(str[:xl-2])
		var sinaks []SinaKline
		dataStr = strings.Replace(dataStr, "d:", "\"d\":", -1)
		dataStr = strings.Replace(dataStr, "o:", "\"o\":", -1)
		dataStr = strings.Replace(dataStr, "h:", "\"h\":", -1)
		dataStr = strings.Replace(dataStr, "l:", "\"l\":", -1)
		dataStr = strings.Replace(dataStr, "c:", "\"c\":", -1)
		dataStr = strings.Replace(dataStr, "v:", "\"v\":", -1)
		err = json.Unmarshal([]byte(dataStr), &sinaks)
		if err != nil {
			log.Println(err)
		}
		for i := range sinaks {
			v := sinaks[i]
			var kx pb.Kline
			if isDaily {
				tm, err := time.ParseInLocation("2006-01-02", v.Day, time.Local)
				if err == nil {
					kx.Time = tm.Unix() * 1000
				}
			} else {
				t, err := time.ParseInLocation("2006-01-02 15:04:05", v.Day, time.Local)
				if err == nil {
					kx.Time = t.Unix() * 1000
				}
			}
			kx.Close, _ = strconv.ParseFloat(v.ClosePrice, 64)
			kx.Open, _ = strconv.ParseFloat(v.OpenPrice, 64)
			kx.High, _ = strconv.ParseFloat(v.MaxPrice, 64)
			kx.Low, _ = strconv.ParseFloat(v.MinPrice, 64)
			kx.Volume, _ = strconv.ParseFloat(v.NowVolume, 64)
			ret.List = append(ret.List, kx)
		}
	}
	return &ret, nil
}

func getOptionSSEKData(symbol *pb.Symbol, period pb.PeriodType, startTime, endTime int64, retryCount int) (*pb.KlineSeries, error) {
	var ret pb.KlineSeries
	// fmt.Println("getOptionSSEKData")
	type SinaKline struct {
		ClosePrice string `json:"c"`
		Day        string `json:"d"`
		MaxPrice   string `json:"h"`
		MinPrice   string `json:"l"`
		NowVolume  string `json:"v"`
		OpenPrice  string `json:"o"`
	}
	url := "http://stock.finance.sina.com.cn/futures/api/jsonp_v2.php/var%20_CON_OP_100014052018_7_4=/StockOptionDaylineService.getSymbolInfo?symbol=" + symbol.Code

	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		tickArr := strings.Split(string(body), "=")
		// log.Printf("------------------")
		// log.Printf(string(tickArr[1]))
		isDaily := true
		xl := len(tickArr[1])
		if xl > 2 && err == nil {
			dataStr := string(tickArr[1][1 : xl-2])
			var sinaks []SinaKline
			dataStr = strings.Replace(dataStr, "d:", "\"d\":", -1)
			dataStr = strings.Replace(dataStr, "o:", "\"o\":", -1)
			dataStr = strings.Replace(dataStr, "h:", "\"h\":", -1)
			dataStr = strings.Replace(dataStr, "l:", "\"l\":", -1)
			dataStr = strings.Replace(dataStr, "c:", "\"c\":", -1)
			dataStr = strings.Replace(dataStr, "v:", "\"v\":", -1)
			err = json.Unmarshal([]byte(dataStr), &sinaks)
			// fmt.Println(err)
			for i := len(sinaks) - 1; i >= 0; i-- {
				v := sinaks[i]
				var kx pb.Kline
				// day := strings.Split(v.Day, " ")[0]
				if isDaily {
					tm, err := time.ParseInLocation("2006-01-02", v.Day, time.Local)
					if err == nil {
						kx.Time = tm.Unix() * 1000
					}
				} else {
					t, err := time.ParseInLocation("2006-01-02 15:04:05", v.Day, time.Local)
					if err == nil {
						kx.Time = t.Unix() * 1000
					}
				}
				v.Day = strings.Split(v.Day, " ")[0]
				kx.Close, _ = strconv.ParseFloat(v.ClosePrice, 64)
				kx.Open, _ = strconv.ParseFloat(v.OpenPrice, 64)
				kx.High, _ = strconv.ParseFloat(v.MaxPrice, 64)
				kx.Low, _ = strconv.ParseFloat(v.MinPrice, 64)
				kx.Volume, _ = strconv.ParseFloat(v.NowVolume, 64)
				ret.List = append(ret.List, kx)
			}
		}
	}
	return &ret, nil
}

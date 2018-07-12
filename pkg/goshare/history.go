package goshare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mineralres/goshare/pkg/base"

	"github.com/mineralres/goshare/pkg/pb"
)

// GetKData 请求历史K线数据
/*
symbol：股票代码，即6位数字代码，或者指数代码
startTime：开始时间time_t
endTime：结束时间time_t
period：周期
retryCount：当网络异常后重试次数，默认为3
*/
func (p *Service) GetKData(symbol *pb.Symbol, period pb.PeriodType, startTime, endTime int64, retryCount int) (*pb.KlineSeries, error) {
	ex := symbol.Exchange
	if ex == pb.ExchangeType_SSE || ex == pb.ExchangeType_SZE {
		return getCNStockKData(symbol, period, startTime, endTime, retryCount)
	} else if ex == pb.ExchangeType_SHFE || ex == pb.ExchangeType_CZCE || ex == pb.ExchangeType_DCE || ex == pb.ExchangeType_CFFEX || ex == pb.ExchangeType_INE {
		return getCNFutureKData(symbol, period, startTime, endTime, retryCount)
	} else if ex == pb.ExchangeType_OPTION_SSE {
		return getOptionSSEKData(symbol, period, startTime, endTime, retryCount)
	}
	var ret pb.KlineSeries
	return &ret, base.ErrUnsupported
}

func getCNStockKData(symbol *pb.Symbol, period pb.PeriodType, startTime, endTime int64, retryCount int) (*pb.KlineSeries, error) {
	var ret pb.KlineSeries
	et := 1
	if symbol.Exchange == pb.ExchangeType_SZE {
		et = 2
	}
	address := fmt.Sprintf("http://pdfm.eastmoney.com/EM_UBG_PDTI_Fast/api/js?rtntype=5&id=%s%d&type=k", symbol.Code, et)
	resp, err := http.Get(address)
	if err != nil {
		return &ret, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &ret, err
	}
	data := string(body)
	data = strings.TrimLeft(data, "(")
	data = strings.TrimRight(data, ")")
	var rtn struct {
		Name string `json:"name"`
		Code string `json:"code"`
		Info struct {
			C string `json:"c"`
		} `json:"info"`
		Data []string `json:"data"`
	}
	err = json.Unmarshal([]byte(data), &rtn)
	if err != nil {
		return &ret, err
	}
	for i := range rtn.Data {
		items := strings.Split(rtn.Data[i], ",")
		if len(items) == 8 {
			var k pb.Kline
			tm, err := time.Parse("2006-01-02", items[0])
			if err != nil {
				log.Println(err, items[0])
				continue
			}
			k.Time = tm.Unix()
			k.Open = base.ParseFloat(items[1])
			k.Close = base.ParseFloat(items[2])
			k.High = base.ParseFloat(items[3])
			k.Low = base.ParseFloat(items[4])
			k.Volume = base.ParseFloat(items[5])
			if strings.Contains(items[6], "万") {
				val := strings.Replace(items[6], "万", "", -1)
				k.Amount = base.ParseFloat(val) * 10000
			} else if strings.Contains(items[6], "亿") {
				val := strings.Replace(items[6], "亿", "", -1)
				k.Amount = base.ParseFloat(val) * 100000000
			} else {
				k.Amount = base.ParseFloat(items[6])
			}
			ret.List = append(ret.List, k)
		}
	}
	log.Println(ret, err)
	return &ret, nil
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
	log.Println(qapi)
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
				tm, err := time.Parse("2006-01-02", v.Day)
				if err == nil {
					kx.Time = (tm.Unix() - 8*3600) * 1000
				}
			} else {
				t, err := time.Parse("2006-01-02 15:04:05", v.Day)
				if err == nil {
					kx.Time = (t.Unix() - 8*3600) * 1000
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
					tm, err := time.Parse("2006-01-02", v.Day)
					if err == nil {
						kx.Time = tm.Unix() * 1000
					}
				} else {
					t, err := time.Parse("2006-01-02 15:04:05", v.Day)
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

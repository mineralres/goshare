package goshare

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"fmt"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/mineralres/goshare/pkg/base"

	"github.com/mineralres/goshare/aproto"
)

// GetLastTick 取最新行情
func (p *Service) GetLastTick(symbol *aproto.Symbol) (*aproto.MarketDataSnapshot, error) {
	if symbol.Exchange == aproto.ExchangeType_SSE || symbol.Exchange == aproto.ExchangeType_SZE {
		return getStockLastTick(symbol)
	}

	if symbol.Exchange == aproto.ExchangeType_INDEX {
		return getIndexLastTick(symbol)
	}

	if symbol.Exchange == aproto.ExchangeType_OPTION_SSE {
		return getOptionSSETick(symbol)
	}

	return nil, base.ErrUnsupported
}

func getStockLastTick(symbol *aproto.Symbol) (*aproto.MarketDataSnapshot, error) {
	ret := &aproto.MarketDataSnapshot{}
	exstr := "sh"
	if symbol.Exchange == aproto.ExchangeType_SZE {
		exstr = "sz"
	}
	tickArr := getRawTickString(exstr, symbol.Code)
	if tickArr == nil || len(tickArr) < 38 {
		return nil, errors.New("ErrGetStockTick")
	}
	if tickArr != nil && len(tickArr) >= 38 {
		timeStr := tickArr[30]
		tx, err := time.Parse("20060102150405", timeStr)
		if err == nil {
			ret.Time = tx.Unix()
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
		var ob5 aproto.OrderBook
		ob5.BidVolume = base.ParseFloat(tickArr[18])
		ob5.Bid = base.ParseFloat(tickArr[17])
		ob5.AskVolume = base.ParseFloat(tickArr[28])
		ob5.Ask = base.ParseFloat(tickArr[27])
		var ob4 aproto.OrderBook
		ob4.BidVolume = base.ParseFloat(tickArr[16])
		ob4.Bid = base.ParseFloat(tickArr[15])
		ob4.AskVolume = base.ParseFloat(tickArr[26])
		ob4.Ask = base.ParseFloat(tickArr[25])
		var ob3 aproto.OrderBook
		ob3.BidVolume = base.ParseFloat(tickArr[14])
		ob3.Bid = base.ParseFloat(tickArr[13])
		ob3.AskVolume = base.ParseFloat(tickArr[24])
		ob3.Ask = base.ParseFloat(tickArr[23])
		var ob2 aproto.OrderBook
		ob2.BidVolume = base.ParseFloat(tickArr[12])
		ob2.Bid = base.ParseFloat(tickArr[11])
		ob2.AskVolume = base.ParseFloat(tickArr[22])
		ob2.Ask = base.ParseFloat(tickArr[21])
		var ob1 aproto.OrderBook
		ob1.BidVolume = base.ParseFloat(tickArr[10])
		ob1.Bid = base.ParseFloat(tickArr[9])
		ob1.AskVolume = base.ParseFloat(tickArr[20])
		ob1.Ask = base.ParseFloat(tickArr[19])
		ret.OrderBookList = []aproto.OrderBook{ob1, ob2, ob3, ob4, ob5}
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
				data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(tickArr[1])), simplifiedchinese.GBK.NewDecoder()))
				if err == nil {
					tickArr[1] = string(data)
				}
			}
			return tickArr
		}
	}
	return nil
}

func getIndexLastTick(symbol *aproto.Symbol) (*aproto.MarketDataSnapshot, error) {
	ret := &aproto.MarketDataSnapshot{}
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

func getOptionSSETick(symbol *aproto.Symbol) (*aproto.MarketDataSnapshot, error) {
	ret := &aproto.MarketDataSnapshot{}
	resp, err := http.Get("http://hq.sinajs.cn/list=" + symbol.Code)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		tickArr := strings.Split(string(body), ",")
		if err == nil && len(tickArr) >= 42 {
			data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(tickArr[37])), simplifiedchinese.GBK.NewDecoder()))
			if err == nil {
				symbol.Code = string(data)
			}
			ret.Symbol = *symbol
			ret.Price = base.ParseFloat(tickArr[2])
			ret.Close = ret.Price
			//ret.PreClose = base.ParseFloat(tickArr[4])
			ret.Open = base.ParseFloat(tickArr[9])
			ret.High = base.ParseFloat(tickArr[39])
			ret.Low = base.ParseFloat(tickArr[40])
			ret.Volume = (base.ParseFloat(tickArr[41]))
			ret.Amount = float64(base.ParseInt(tickArr[42]))
			ret.UpperLimitPrice = base.ParseFloat(tickArr[10])
			ret.LowerLimitPrice = base.ParseFloat(tickArr[11])
			var ob5 aproto.OrderBook
			ob5.BidVolume = base.ParseFloat(tickArr[12])
			ob5.Bid = base.ParseFloat(tickArr[13])
			ob5.AskVolume = base.ParseFloat(tickArr[30])
			ob5.Ask = base.ParseFloat(tickArr[31])
			var ob4 aproto.OrderBook
			ob4.BidVolume = base.ParseFloat(tickArr[14])
			ob4.Bid = base.ParseFloat(tickArr[15])
			ob4.AskVolume = base.ParseFloat(tickArr[28])
			ob4.Ask = base.ParseFloat(tickArr[29])
			var ob3 aproto.OrderBook
			ob3.BidVolume = base.ParseFloat(tickArr[16])
			ob3.Bid = base.ParseFloat(tickArr[17])
			ob3.AskVolume = base.ParseFloat(tickArr[26])
			ob3.Ask = base.ParseFloat(tickArr[27])
			var ob2 aproto.OrderBook
			ob2.BidVolume = base.ParseFloat(tickArr[18])
			ob2.Bid = base.ParseFloat(tickArr[19])
			ob2.AskVolume = base.ParseFloat(tickArr[24])
			ob2.Ask = base.ParseFloat(tickArr[25])
			var ob1 aproto.OrderBook
			ob1.BidVolume = base.ParseFloat(tickArr[20])
			ob1.Bid = base.ParseFloat(tickArr[21])
			ob1.AskVolume = base.ParseFloat(tickArr[22])
			ob1.Ask = base.ParseFloat(tickArr[23])
			//Log(symbol.Code)
			return ret, nil
		}
	}
	return nil, errors.New("ErrGetIndex")
}

// 统一日志打印
func Log(sd string) {
	log.Printf(sd)
}

// GetSina50EtfSym 获取50ETF期权合约列表，sina代码
func (p *Service) GetSina50EtfSym(sym string) (slice []string) {
	resp, err := http.Get("http://hq.sinajs.cn/list=" + sym)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		tickArr := strings.Split(string(body), ",")
		i := len(tickArr)
		if err == nil {
			slice = make([]string, i-2)
			for j := 1; j < i-1; j++ {
				slice[j-1] = tickArr[j]
			}
			return slice
		}
	}
	return nil
}


// GetMainFutureLastTick 
func (p *Service) GetMainFutureLastTick(et aproto.ExchangeType) ([]aproto.MarketDataSnapshot, error) {
	var ret []aproto.MarketDataSnapshot
	// address := fmt.Sprintf("http://nufm.dfcfw.com/EM_Finance2014NumericApplication/JS.aspx?type=ct&st=(BalFlowMain)&sr=-1&p=1&ps=%d", size) + "&js=var%20PPHMDFMQ={pages:(pc),date:%222014-10-22%22,data:[(x)]}&token=894050c76af8597a853f5b408b759f5d&cmd=C._AB&sty=DCFFITA&rt=50714413"
	var et_str string;	
	switch et {
	case aproto.ExchangeType_SHFE:
		et_str = "SHFE"
	case aproto.ExchangeType_CZCE:
		et_str = "CZCE"
	case aproto.ExchangeType_DCE:
		et_str = "DCE"
	case aproto.ExchangeType_CFFEX:
		et_str = "_168"
	default:
		return ret, fmt.Errorf("error ExchangeType %s", et)
	}

	address := fmt.Sprintf("http://nufm.dfcfw.com/EM_Finance2014NumericApplication/JS.aspx?type=CT&cmd=C.%s", et_str) + "&sty=FCFL4O&sortType=(ChangePercent)&sortRule=-1&page=1&pageSize=200&js={rank:[(x)],pages:(pc),total:(tot)}&token=7bc05d0d4c3c22ef9fca8c2a912d779c&jsName=quote_123&_g=0.628606915911589&_=1521620666159";

	resp, err := http.Get(address)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err == nil {
			tickArr := strings.Split(string(body), "\"")
			// fmt.Println(string(body))
			i := 0
			for i < len(tickArr){

				mkt_str_arr := strings.Split(string(tickArr[i]), ",")
				i++
				if len(mkt_str_arr) < 15{
					continue
				}
				if len(mkt_str_arr[1]) > 3{
					continue;
				}

				mkt := aproto.MarketDataSnapshot{}
				mkt.Symbol = aproto.Symbol{Exchange: et, Code: mkt_str_arr[1]}
				mkt.Open = base.ParseFloat(mkt_str_arr[11])
				mkt.High = base.ParseFloat(mkt_str_arr[13])
				mkt.Low = base.ParseFloat(mkt_str_arr[14])
				mkt.Price = base.ParseFloat(mkt_str_arr[3])
				mkt.Close = mkt.Price
				mkt.Volume = base.ParseFloat(mkt_str_arr[10])
				mkt.Amount = base.ParseFloat(mkt_str_arr[15])
				mkt.Position = base.ParseFloat(mkt_str_arr[9])
				mkt.PreSettlementPrice = base.ParseFloat(mkt_str_arr[8])
				ret = append(ret, mkt)
				// fmt.Println(mkt.Symbol, mkt.Open, mkt.High, mkt.Low)
			}
		}

	}
	return ret, fmt.Errorf("error ExchangeType %s", et)
}
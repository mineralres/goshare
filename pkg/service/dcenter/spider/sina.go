package spider

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	pb "github.com/mineralres/goshare/pkg/pb/goshare"
	"github.com/mineralres/goshare/pkg/util"
)

// BatchGetSSEStockOptionTick 取所有行情
func (s *Spider) BatchGetSSEStockOptionTick(symbols []pb.Symbol) ([]pb.MarketDataSnapshot, error) {
	rets := []pb.MarketDataSnapshot{}
	all := "http://hq.sinajs.cn/list="
	for i := range symbols {
		all = all + "CON_OP_" + symbols[i].Code + ","
	}
	resp, err := http.Get(all)
	if err != nil {
		return nil, errors.New("获取原始数据错误")
	}

	defer resp.Body.Close()
	bodyx, _ := ioutil.ReadAll(resp.Body)
	str := util.Decode(string(bodyx))
	tickArr1 := strings.Split(str, ";")
	for _, v := range tickArr1 {
		tickArr := strings.Split(v, ",")
		ret := pb.MarketDataSnapshot{}
		if err == nil && len(tickArr) >= 42 {
			pos := strings.Index(tickArr[0], "1000")
			symbol := &pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: tickArr[0][pos : pos+8]}
			ret.Symbol = symbol
			ret.Price = util.ParseFloat(tickArr[2])
			ret.Close = ret.Price
			ret.Position = int32(util.ParseInt(tickArr[5]))
			ret.Open = util.ParseFloat(tickArr[9])
			ret.High = util.ParseFloat(tickArr[39])
			ret.Low = util.ParseFloat(tickArr[40])
			ret.Volume = int32(util.ParseInt(tickArr[41]))
			ret.Amount = float64(util.ParseInt(tickArr[42]))
			ret.UpperLimitPrice = util.ParseFloat(tickArr[10])
			ret.LowerLimitPrice = util.ParseFloat(tickArr[11])
			ret.ExercisePrice = util.ParseFloat(tickArr[7])
			ret.PreClose = util.ParseFloat(tickArr[8])

			var ob5 pb.OrderBook
			ob5.Ask = util.ParseFloat(tickArr[12])
			ob5.AskVolume = int32(util.ParseInt(tickArr[13]))
			ob5.Bid = util.ParseFloat(tickArr[30])
			ob5.BidVolume = int32(util.ParseInt(tickArr[31]))
			var ob4 pb.OrderBook
			ob4.Ask = util.ParseFloat(tickArr[14])
			ob4.AskVolume = int32(util.ParseInt(tickArr[15]))
			ob4.Bid = util.ParseFloat(tickArr[28])
			ob4.BidVolume = int32(util.ParseInt(tickArr[29]))
			var ob3 pb.OrderBook
			ob3.Ask = util.ParseFloat(tickArr[16])
			ob3.AskVolume = int32(util.ParseInt(tickArr[17]))
			ob3.Bid = util.ParseFloat(tickArr[26])
			ob3.BidVolume = int32(util.ParseInt(tickArr[27]))
			var ob2 pb.OrderBook
			ob2.Ask = util.ParseFloat(tickArr[18])
			ob2.AskVolume = int32(util.ParseInt(tickArr[19]))
			ob2.Bid = util.ParseFloat(tickArr[24])
			ob2.BidVolume = int32(util.ParseInt(tickArr[25]))
			var ob1 pb.OrderBook
			ob1.Ask = util.ParseFloat(tickArr[20])
			ob1.AskVolume = int32(util.ParseInt(tickArr[21]))
			ob1.Bid = util.ParseFloat(tickArr[22])
			ob1.BidVolume = int32(util.ParseInt(tickArr[23]))
			ret.OrderBookList = []*pb.OrderBook{&ob1, &ob2, &ob3, &ob4, &ob5}
			ret.Name = tickArr[37]

			ret.Time = util.ParseBeijingTime("2006-01-02 15:04:05", tickArr[32])
			td, err := strconv.Atoi(time.Unix(ret.Time, 0).Format("20060102"))
			if err == nil {
				ret.TradingDay = int32(td)
			}

			rets = append(rets, ret)
		}
	}
	return rets, nil
}

// GetOptionSinaTick 根据交割月获取t型报价表数据
/* date 如1808 为8月到期的
 */
func (s *Spider) GetOptionSinaTick(date string) ([]pb.MarketDataSnapshot, error) {
	rets := []pb.MarketDataSnapshot{}

	all := "OP_DOWN_510050" + date
	allTick, _, _ := getOptionSSETickT(all)
	rets = append(rets, allTick...)

	all = "OP_UP_510050" + date
	allTick, _, _ = getOptionSSETickT(all)
	rets = append(rets, allTick...)

	return rets, nil
}

// GetOptionTQuote 根据交割月获取t型报价表数据
/* date 如1808 为8月到期的
 */
func (s *Spider) GetOptionTQuote(date string) ([]pb.OptionTMarket, error) {
	rets := []pb.OptionTMarket{}

	all := "OP_DOWN_510050" + date
	allTick, allName, _ := getOptionSSETickT(all)

	all = "OP_UP_510050" + date
	allTick1, _, _ := getOptionSSETickT(all)

	for kk := range allName {
		msg := pb.OptionTMarket{}
		msg.CallTk = &allTick1[kk]
		msg.PutTk = &allTick[kk]
		rets = append(rets, msg)
		//log.Printf("执行价为%s,call 为%s,put 为%s", val, msg.CallTk.Symbol.Code, msg.PutTk.Symbol.Code)
	}
	return rets, nil
}

// GetSina50EtfSym 获取50ETF期权合约列表，sina代码
//说明：
//OP_DOWN_5100501807:OP 期权、DOWN 看跌、UP 看涨、510050 50etf标的代码、1807 到期月份
//根据到期月的期权从接口获取t型的合约表： CON_OP_10001394
// 参数解释：CON_OP_ 为固定title，10001394这个是交易所的合约代码，在任何一个行情软件都可以查到，也可以通过GetSina50EtfSym接口获取
// GetLastTick 根据CON_OP_10001394可以获取最新的报价
// GetKData 根据CON_OP_10001394可以获取日k线
func GetSina50EtfSym(sym string) []string {
	var ret []string
	resp, err := http.Get("http://hq.sinajs.cn/list=" + sym)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ret
		}
		tickArr := strings.Split(string(body), "=")
		if len(tickArr) != 2 {
			return ret
		}
		str := strings.TrimLeft(tickArr[1], "\"")
		tickArr = strings.Split(str, ",")
		// log.Println("tickArr", tickArr, str)
		for i := range tickArr {
			if len(tickArr[i]) > 3 {
				ret = append(ret, tickArr[i])
			}
		}
	}
	return ret
}

// parse sina tick string
func parseSinaOptionTick(body string) (*pb.MarketDataSnapshot, string, error) {
	ret := &pb.MarketDataSnapshot{}
	tickArr := strings.Split(string(body), ",")
	if len(tickArr) >= 42 {
		var ss string
		tickSym2 := strings.Split(strings.Split(tickArr[0], "=")[0], "_")
		ss = tickSym2[4]
		symbol := &pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: ss}
		ret.Symbol = symbol
		ret.Price = util.ParseFloat(tickArr[2])
		ret.Close = ret.Price
		ret.ExercisePrice = util.ParseFloat(tickArr[7])
		ret.PreClose = util.ParseFloat(tickArr[8])
		ret.Open = util.ParseFloat(tickArr[9])
		ret.High = util.ParseFloat(tickArr[39])
		ret.Low = util.ParseFloat(tickArr[40])
		ret.Volume = int32(util.ParseInt(tickArr[41]))
		ret.Amount = (util.ParseFloat(tickArr[42]))
		ret.UpperLimitPrice = util.ParseFloat(tickArr[10])
		ret.LowerLimitPrice = util.ParseFloat(tickArr[11])
		var ob5 pb.OrderBook
		ob5.Ask = util.ParseFloat(tickArr[12])
		ob5.AskVolume = int32(util.ParseInt(tickArr[13]))
		ob5.Bid = util.ParseFloat(tickArr[30])
		ob5.BidVolume = int32(util.ParseInt(tickArr[31]))
		var ob4 pb.OrderBook
		ob4.Ask = util.ParseFloat(tickArr[14])
		ob4.AskVolume = int32(util.ParseInt(tickArr[15]))
		ob4.Bid = util.ParseFloat(tickArr[28])
		ob4.BidVolume = int32(util.ParseInt(tickArr[29]))
		var ob3 pb.OrderBook
		ob3.Ask = util.ParseFloat(tickArr[16])
		ob3.AskVolume = int32(util.ParseInt(tickArr[17]))
		ob3.Bid = util.ParseFloat(tickArr[26])
		ob3.BidVolume = int32(util.ParseInt(tickArr[27]))
		var ob2 pb.OrderBook
		ob2.Ask = util.ParseFloat(tickArr[18])
		ob2.AskVolume = int32(util.ParseInt(tickArr[19]))
		ob2.Bid = util.ParseFloat(tickArr[24])
		ob2.BidVolume = int32(util.ParseInt(tickArr[25]))
		var ob1 pb.OrderBook
		ob1.Ask = util.ParseFloat(tickArr[20])
		ob1.AskVolume = int32(util.ParseInt(tickArr[21]))
		ob1.Bid = util.ParseFloat(tickArr[22])
		ob1.BidVolume = int32(util.ParseInt(tickArr[23]))
		ret.OrderBookList = []*pb.OrderBook{&ob1, &ob2, &ob3, &ob4, &ob5}
		ret.Name = util.StringFromGBK(tickArr[37])
		// timex, err := time.Parse("2006-01-02 15:04:05", tickArr[32])
		ret.Time = util.ParseBeijingTime("2006-01-02 15:04:05", tickArr[32])
		return ret, tickArr[37], nil
	}
	return nil, "", errors.New("error")
}

// getSSEOptionTick 根据合约获取单个期权合约的tick数据
func getSSEOptionTick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	//ret := &pb.MarketDataSnapshot{}
	resp, err := http.Get("http://hq.sinajs.cn/list=CON_OP_" + symbol.Code)
	if err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		ret, _, err1 := parseSinaOptionTick(string(body))
		if err1 == nil {
			td, err := strconv.Atoi(time.Unix(ret.Time, 0).Format("20060102"))
			if err == nil {
				ret.TradingDay = int32(td)
			}
			return ret, nil
		}
	}
	return nil, errors.New("ErrGetIndex")
}

// 批量获取50etf tick数据
func getOptionSSETickT(symbol string) ([]pb.MarketDataSnapshot, []string, error) {
	rets := []pb.MarketDataSnapshot{}
	retsName := []string{}
	syms := GetSina50EtfSym(symbol)
	all := "http://hq.sinajs.cn/list="
	for _, value := range syms {
		all = all + value + ","
	}
	// log.Printf(" sina 期权合约代码为: %s\n", all)
	resp, err := http.Get(all)
	if err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		tickArr1 := strings.Split(string(body), ";")
		for _, v := range tickArr1 {
			ret, name, err1 := parseSinaOptionTick(string(v))
			if err1 == nil {
				rets = append(rets, *ret)
				retsName = append(retsName, name)
			}
		}
		return rets, retsName, nil
	}
	return nil, nil, errors.New("ErrGetIndex")
}

// GetSSEStockOptionTick 取所有行情
func (s *Spider) GetSSEStockOptionTick(symbols []pb.Symbol) ([]pb.MarketDataSnapshot, error) {
	rets := []pb.MarketDataSnapshot{}
	all := "http://hq.sinajs.cn/list="
	for _, value := range symbols {
		all = all + "CON_OP_" + value.Code + ","
	}
	resp, err := http.Get(all)
	if err != nil {
		return nil, errors.New("ErrGetIndex")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	tickArr1 := strings.Split(string(body), ";")
	for _, v := range tickArr1 {
		tickArr := strings.Split(v, ",")
		ret := pb.MarketDataSnapshot{}
		if err == nil && len(tickArr) >= 42 {
			symbol := &pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: tickArr[0][19:27]}
			ret.Symbol = symbol
			ret.Price = util.ParseFloat(tickArr[2])
			ret.Close = ret.Price
			ret.Position = util.ParseInt32(tickArr[5])
			ret.Open = util.ParseFloat(tickArr[9])
			ret.High = util.ParseFloat(tickArr[39])
			ret.Low = util.ParseFloat(tickArr[40])
			ret.Volume = (util.ParseInt32(tickArr[41]))
			ret.Amount = float64(util.ParseInt(tickArr[42]))
			ret.UpperLimitPrice = util.ParseFloat(tickArr[10])
			ret.LowerLimitPrice = util.ParseFloat(tickArr[11])
			var ob5 pb.OrderBook
			ob5.BidVolume = util.ParseInt32(tickArr[12])
			ob5.Bid = util.ParseFloat(tickArr[13])
			ob5.AskVolume = util.ParseInt32(tickArr[30])
			ob5.Ask = util.ParseFloat(tickArr[31])
			var ob4 pb.OrderBook
			ob4.BidVolume = util.ParseInt32(tickArr[14])
			ob4.Bid = util.ParseFloat(tickArr[15])
			ob4.AskVolume = util.ParseInt32(tickArr[28])
			ob4.Ask = util.ParseFloat(tickArr[29])
			var ob3 pb.OrderBook
			ob3.BidVolume = util.ParseInt32(tickArr[16])
			ob3.Bid = util.ParseFloat(tickArr[17])
			ob3.AskVolume = util.ParseInt32(tickArr[26])
			ob3.Ask = util.ParseFloat(tickArr[27])
			var ob2 pb.OrderBook
			ob2.BidVolume = util.ParseInt32(tickArr[18])
			ob2.Bid = util.ParseFloat(tickArr[19])
			ob2.AskVolume = util.ParseInt32(tickArr[24])
			ob2.Ask = util.ParseFloat(tickArr[25])
			var ob1 pb.OrderBook
			ob1.BidVolume = util.ParseInt32(tickArr[20])
			ob1.Bid = util.ParseFloat(tickArr[21])
			ob1.AskVolume = util.ParseInt32(tickArr[22])
			ob1.Ask = util.ParseFloat(tickArr[23])
			rets = append(rets, ret)
		}
	}
	return rets, nil
}

func getRawTickString(exstr string, symbol string) []string {
	resp, err := http.Get("http://web.sqt.gtimg.cn/q=" + exstr + symbol)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			tickArr := strings.Split(string(body), "~")
			if len(tickArr) > 1 {
				tickArr[1] = util.StringFromGBK(tickArr[1])
			} else {
				log.Printf("getRawTickString %s-%s", exstr, symbol)
			}
			return tickArr
		}
	}
	return nil
}

func parseSinaTime(layout, value string) int64 {
	return util.ParseBeijingTime(layout, value)
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
		ret.Time = parseSinaTime("20060102150405", timeStr)
		td, err := strconv.Atoi(time.Unix(ret.Time, 0).Format("20060102"))
		if err == nil {
			ret.TradingDay = int32(td)
		}

		ret.Symbol = symbol
		ret.Price = util.ParseFloat(tickArr[3])
		ret.Close = ret.Price
		ret.PreClose = util.ParseFloat(tickArr[4])
		ret.Open = util.ParseFloat(tickArr[5])
		ret.High = util.ParseFloat(tickArr[33])
		ret.Low = util.ParseFloat(tickArr[34])
		ret.Volume = (util.ParseInt32(tickArr[6]))
		ret.Amount = float64(util.ParseInt(tickArr[37]) * 10000)
		ret.UpperLimitPrice = util.ParseFloat(tickArr[47])
		ret.LowerLimitPrice = util.ParseFloat(tickArr[48])
		var ob5 pb.OrderBook
		ob5.BidVolume = util.ParseInt32(tickArr[18])
		ob5.Bid = util.ParseFloat(tickArr[17])
		ob5.AskVolume = util.ParseInt32(tickArr[28])
		ob5.Ask = util.ParseFloat(tickArr[27])
		var ob4 pb.OrderBook
		ob4.BidVolume = util.ParseInt32(tickArr[16])
		ob4.Bid = util.ParseFloat(tickArr[15])
		ob4.AskVolume = util.ParseInt32(tickArr[26])
		ob4.Ask = util.ParseFloat(tickArr[25])
		var ob3 pb.OrderBook
		ob3.BidVolume = util.ParseInt32(tickArr[14])
		ob3.Bid = util.ParseFloat(tickArr[13])
		ob3.AskVolume = util.ParseInt32(tickArr[24])
		ob3.Ask = util.ParseFloat(tickArr[23])
		var ob2 pb.OrderBook
		ob2.BidVolume = util.ParseInt32(tickArr[12])
		ob2.Bid = util.ParseFloat(tickArr[11])
		ob2.AskVolume = util.ParseInt32(tickArr[22])
		ob2.Ask = util.ParseFloat(tickArr[21])
		var ob1 pb.OrderBook
		ob1.BidVolume = util.ParseInt32(tickArr[10])
		ob1.Bid = util.ParseFloat(tickArr[9])
		ob1.AskVolume = util.ParseInt32(tickArr[20])
		ob1.Ask = util.ParseFloat(tickArr[19])
		ret.OrderBookList = []*pb.OrderBook{&ob1, &ob2, &ob3, &ob4, &ob5}
	}
	return ret, nil
}

// GetIndexLastTick 指数行情
func (s *Spider) GetIndexLastTick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	ret := &pb.MarketDataSnapshot{}
	resp, err := http.Get("http://hq.sinajs.cn/list=" + symbol.Code)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			tickArr := strings.Split(string(body), ",")
			//sym := strings.Split(string(tickArr[0]), "=")
			ret.Symbol = symbol
			ret.Price = util.ParseFloat(tickArr[1])
			ret.Close = ret.Price
			return ret, nil
		}
	}
	return nil, errors.New("ErrGetIndex")
}

// GetLastTick 获取最新报价
func (s *Spider) GetLastTick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
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
		return s.GetIndexLastTick(symbol)
	}
	return nil, errors.New("unsported")
}

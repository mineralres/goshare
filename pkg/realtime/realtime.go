package realtime

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/mineralres/goshare/pkg/base"

	"github.com/mineralres/goshare/aproto"
)

// RProvider RProvider
type RProvider struct {
}

// GetLastTick 取最新行情
func (p *RProvider) GetLastTick(symbol *aproto.Symbol) (*aproto.MarketDataSnapshot, error) {
	if symbol.Exchange == aproto.ExchangeType_SSE || symbol.Exchange == aproto.ExchangeType_SZE {
		return getStockLastTick(symbol)
	}
	return nil, errors.New("不支持的交易所类型")
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

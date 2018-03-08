package history

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mineralres/goshare/pkg/base"

	"github.com/mineralres/goshare/aproto"
)

// GetKData 请求历史K线数据
/*
symbol：股票代码，即6位数字代码，或者指数代码（sh=上证指数 sz=深圳成指 hs300=沪深300指数 sz50=上证50 zxb=中小板 cyb=创业板）
startDate：开始日期，格式20180307
endDate：结束日期，格式20180307
period：周期
retryCount：当网络异常后重试次数，默认为3
*/
func GetKData(symbol *aproto.Symbol, period aproto.PeriodType, startDate, endDate, retryCount int) (*aproto.KlineSeries, error) {
	if symbol.Exchange == aproto.ExchangeType_SSE || symbol.Exchange == aproto.ExchangeType_SZE {
		return getCNStockKData(symbol, period, startDate, endDate, retryCount)
	}
	var ret aproto.KlineSeries
	return &ret, nil
}

func getCNStockKData(symbol *aproto.Symbol, period aproto.PeriodType, startDate, endDate, retryCount int) (*aproto.KlineSeries, error) {
	var ret aproto.KlineSeries
	et := 1
	if symbol.Exchange == aproto.ExchangeType_SZE {
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
			var k aproto.Kline
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

package goshare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mineralres/goshare/pkg/base"
	"github.com/mineralres/goshare/pkg/pb"
)

// EaseMoneySource 东方财富
type EaseMoneySource struct {
}

// GetRealtimeMoneyTrendList 取实时资金流向. size 需要取的条数
func (s *EaseMoneySource) GetRealtimeMoneyTrendList(size int) (*pb.RealtimeMoneyTrendItemList, error) {
	var ret pb.RealtimeMoneyTrendItemList
	address := fmt.Sprintf("http://nufm.dfcfw.com/EM_Finance2014NumericApplication/JS.aspx?type=ct&st=(BalFlowMain)&sr=-1&p=1&ps=%d", size) + "&js=var%20PPHMDFMQ={pages:(pc),date:%222014-10-22%22,data:[(x)]}&token=894050c76af8597a853f5b408b759f5d&cmd=C._AB&sty=DCFFITA&rt=50714413"
	resp, err := http.Get(address)
	if err != nil {
		return &ret, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &ret, err
	}
	content := string(body)
	content = strings.Replace(content, "var PPHMDFMQ=", "", 1)
	content = strings.Replace(content, "pages", `"pages"`, 1)
	content = strings.Replace(content, "date", `"date"`, 1)
	content = strings.Replace(content, "data", `"data"`, 1)
	var rtn struct {
		Pages int      `json:"pages"`
		Date  string   `json:"date"`
		Data  []string `json:"data"`
	}
	err = json.Unmarshal([]byte(content), &rtn)
	if err == nil {
		for i := range rtn.Data {
			str := &rtn.Data[i]
			items := strings.Split(*str, ",")
			if len(items) > 14 {
				var item pb.RealtimeMoneyTrendItem
				item.Symbol.Code = items[1]
				item.Name = items[2]
				item.Price = base.ParseFloat(items[3])
				item.UpdownPercentage = base.ParseFloat(items[4]) / 100
				item.SuperSuperBigOrder.Amount = base.ParseFloat(items[5]) * 10000
				item.SuperSuperBigOrder.Percentage = base.ParseFloat(items[6]) / 100
				item.SuperBigOrder.Amount = base.ParseFloat(items[7]) * 10000
				item.SuperBigOrder.Percentage = base.ParseFloat(items[8]) / 100
				item.BigOrder.Amount = base.ParseFloat(items[9]) * 10000
				item.BigOrder.Percentage = base.ParseFloat(items[10]) / 100
				item.MiddleOrder.Amount = base.ParseFloat(items[11]) * 10000
				item.MiddleOrder.Percentage = base.ParseFloat(items[12]) / 100
				item.SmallOrder.Amount = base.ParseFloat(items[13]) * 10000
				item.SmallOrder.Percentage = base.ParseFloat(items[14]) / 100
				t, err := time.ParseInLocation("2006-01-02 15:04:05", items[15], time.Local)
				if err == nil {
					item.Time = t.Unix()
				}
				ret.List = append(ret.List, item)
			}
		}
	}
	return &ret, nil
}

// GetCNStockKData 股票K线.
func (s *EaseMoneySource) GetCNStockKData(symbol *pb.Symbol, period pb.PeriodType, startTime, endTime int64, retryCount int) (*pb.KlineSeries, error) {
	var ret pb.KlineSeries
	et := 1
	if symbol.Exchange == pb.ExchangeType_SZE {
		et = 2
	}
	ktype := "k" // d1
	if period == pb.PeriodType_M5 {
		ktype = "m5k"
	} else if period == pb.PeriodType_M15 {
		ktype = "m15k"
	} else if period == pb.PeriodType_M1 {
		ktype = "m1k"
	} else if period == pb.PeriodType_M30 {
		ktype = "m30k"
	} else if period == pb.PeriodType_H1 {
		ktype = "m60k"
	}

	authorityType := "fa"

	address := fmt.Sprintf("http://pdfm.eastmoney.com/EM_UBG_PDTI_Fast/api/js?rtntype=5&id=%s%d&type=%s&authorityType=%s", symbol.Code, et, ktype, authorityType)
	// log.Println(address)
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
		if len(items) >= 8 {
			var k pb.Kline
			layoutStr := "2006-01-02 15:04"
			if period == pb.PeriodType_D1 {
				layoutStr = "2006-01-02"
			}
			tm, err := time.ParseInLocation(layoutStr, items[0], time.Local)
			if err != nil {
				log.Println(err, items[0])
				continue
			}
			k.Time = (tm.Unix() - 3600*8) * 1000
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
	return &ret, nil
}

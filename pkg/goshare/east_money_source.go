package goshare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
				t, err := time.Parse("2006-01-02 15:04:05", items[15])
				if err == nil {
					item.Time = t.Unix()
				}
				ret.List = append(ret.List, item)
			}
		}
	}
	return &ret, nil
}

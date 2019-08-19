package spider

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	gspb "github.com/mineralres/goshare/pkg/pb/goshare"
	pb "github.com/mineralres/goshare/pkg/pb/spider"
)

// Xueqiu  xueqiu
type Xueqiu struct {
}

func (xq *Xueqiu) getURLContent(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Cookie", "device_id=23be3ba9d7efa08daedc6b0c8c4217b6; _ga=GA1.2.143596347.1560853613; s=dg142zzimp; xq_a_token=17067303557fc0af0961063ffb2aa2341c3132a4; xq_a_token.sig=pw-CFTAO0pv_iu47dMProDI3rw4; xq_r_token=c1476ba66a6a12fbe62ab833a29e4445bc84385e; xq_r_token.sig=B7iwDRzZ1x5VWNYfZob70xUP71E; u=261564729466741; Hm_lvt_1db88642e346389874251b5a1eded6e3=1564148535,1564148691,1564729467; Hm_lpvt_1db88642e346389874251b5a1eded6e3=1564729845")
	req.Header.Set("Host", "stock.xueqiu.com")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	reader, err := gzip.NewReader(res.Body)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return b, err
}

// HotStockList  hot stock list
// typ 10 1hour , 20 24hours
func (xq *Xueqiu) hotStockList(t1, t2 string) ([]*pb.HotStockItem, error) {
	url := fmt.Sprintf(`https://stock.xueqiu.com/v5/stock/hot_stock/list.json?size=8&_type=%s&type=%s`, t1, t2)
	body, err := xq.getURLContent(url)
	if err != nil {
		return nil, err
	}
	type hotRankItem struct {
		Change    float64 `json:"chg"`
		Code      string  `json:"code"`
		Current   float64 `json:"current"`
		Increment float64 `json:"increment"`
		Name      string  `json:"name"`
		Percent   float64 `json:"percent"`
		Type      int32   `json:"type"`
		Value     float64 `json:"value"`
	}

	var ret struct {
		Data struct {
			Items []*hotRankItem `json:"items"`
		} `json:"data"`
	}
	err = json.Unmarshal(body, &ret)
	var arr []*pb.HotStockItem
	if err == nil {
		for _, item := range ret.Data.Items {
			var h pb.HotStockItem
			h.Exchange = item.Code[:2]
			h.Symbol = item.Code[2:]
			h.Change = item.Change
			h.Current = item.Current
			h.Increment = item.Increment
			h.Name = item.Name
			h.Percent = item.Percent
			h.Type = item.Type
			h.Value = item.Value
			arr = append(arr, &h)
		}
	}
	return arr, err
}

// HotStockList hot stock list
func (xq *Xueqiu) HotStockList() (*pb.HotStockList, error) {
	var ret pb.HotStockList
	var err error
	ret.GlobalH1, err = xq.hotStockList("10", "10")
	if err != nil {
		return &ret, err
	}
	ret.GlobalH24, err = xq.hotStockList("10", "20")
	if err != nil {
		return &ret, err
	}
	ret.AshareH1, err = xq.hotStockList("12", "12")
	if err != nil {
		return &ret, err
	}
	ret.AshareH24, err = xq.hotStockList("12", "22")
	if err != nil {
		return &ret, err
	}
	return &ret, nil
}

// StarCount star count
func (xq *Xueqiu) StarCount(exchange, symbol string) (int, error) {
	if exchange == "SSE" {
		symbol = "SH" + symbol
	} else if exchange == "SZE" {
		symbol = "SZ" + symbol
	} else {
		return 0, errors.New("unsported exchange")
	}
	url := fmt.Sprintf(`https://xueqiu.com/recommend/pofriends.json?type=1&code=%s&start=0&count=14`, symbol)
	body, err := xq.getURLContent(url)
	var res struct {
		TotalCount int `json:"totalcount"`
	}
	err = json.Unmarshal(body, &res)
	return res.TotalCount, err
}

// BonusHistory stock bonus
func (xq *Xueqiu) BonusHistory(exchange, symbol string) ([]*pb.Bonus, error) {
	if exchange == "SSE" {
		exchange = "SH"
	} else if exchange == "SZE" {
		exchange = "SZ"
	} else {
		return nil, errors.New("unsupported exchange")
	}
	url := fmt.Sprintf(`https://stock.xueqiu.com/v5/stock/f10/cn/bonus.json?symbol=%s%s&size=70&page=1&extend=true`, exchange, symbol)
	body, err := xq.getURLContent(url)
	if err != nil {
		return nil, err
	}
	type bonusItem struct {
		Ashare_ex_dividend_date int64  `json:"ashare_ex_dividend_date"`
		Cancle_dividend_date    int64  `json:"cancle_dividend_date"`
		Dividend_date           int64  `json:"dividend_date"`
		Dividend_year           string `json:"dividend_year"`
		Equity_date             int64  `json:"equity_date"`
		Ex_dividend_date        int64  `json:"ex_dividend_date"`
		Plan_explain            string `json:"plan_explain"`
	}
	var ret struct {
		Data struct {
			Items []*bonusItem `json:"items"`
		} `json:"data"`
	}
	err = json.Unmarshal(body, &ret)
	var l []*pb.Bonus
	// v1 := regexp.MustCompile(`派([0-9])元`)
	v2 := regexp.MustCompile(`转([0-9])股`)
	v1 := regexp.MustCompile(`派([0-9]*.[0-9]*)元`)
	for _, i := range ret.Data.Items {
		var b pb.Bonus
		b.PlanExplain = i.Plan_explain
		b.ExDividendDate = i.Ex_dividend_date / 1000
		b.EquityDate = i.Equity_date / 1000
		b.DividendYear = i.Dividend_year
		b.DividendDate = i.Dividend_date / 1000
		b.CancleDividendDate = i.Cancle_dividend_date / 1000
		b.AshareExDividendDate = i.Ashare_ex_dividend_date / 1000
		params := v1.FindStringSubmatch(b.PlanExplain)
		if len(params) == 2 {
			cash, err := strconv.ParseFloat(params[1], 64)
			if err == nil {
				b.DividendCash = cash / 10
			}
		}
		params = v2.FindStringSubmatch(b.PlanExplain)
		if len(params) == 2 {
			share, err := strconv.ParseFloat(params[1], 64)
			if err == nil {
				b.DividendShare = share / 10
			}
		}
		l = append(l, &b)
	}
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
	return l, nil
}

// KlineSeries kline series
// xtype 复制类型 before, after, normal
func (xq *Xueqiu) KlineSeries(exchange, symbol string, period gspb.PeriodType, xtype string, begin, count int64) ([]*gspb.Kline, error) {
	if exchange == "SSE" {
		exchange = "SH"
	} else if exchange == "SZE" {
		exchange = "SZ"
	} else {
		return nil, errors.New("unsupported exchange")
	}
	t := "week"
	switch period {
	case gspb.PeriodType_D1:
		t = "day"
	}
	var arr []*gspb.Kline
	url := fmt.Sprintf(`https://stock.xueqiu.com/v5/stock/chart/kline.json?symbol=%s%s&begin=%d&period=%s&type=%s&count=%d&indicator=kline`, exchange, symbol, begin, t, xtype, count)
	body, err := xq.getURLContent(url)
	if err != nil {
		return nil, err
	}
	var ret struct {
		Data struct {
			Item [][]float64 `json:"item"`
		} `json:"data"`
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return arr, err
	}
	for i := range ret.Data.Item {
		row := ret.Data.Item[i]
		if len(row) < 12 {
			panic("")
		}
		var k gspb.Kline
		k.Time = int64(row[0] / 1000)
		k.Volume = int32((row[1]))
		k.Open = (row[2])
		k.High = (row[3])
		k.Low = (row[4])
		k.Close = (row[5])
		k.Amount = (row[9])
		arr = append(arr, &k)
	}
	return arr, nil
}

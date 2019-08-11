package spider

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/mineralres/protos/src/go/goshare"
	"github.com/mineralres/goshare/pkg/spider"
	"github.com/mineralres/goshare/pkg/tdxclient"
)

type bonusConfig struct {
	CachePath             string `json:"cachePath"`             // 缓存路径 可以删除以获得最新数据
	OutputPath            string `json:"outputPath"`            // 输入文件名
	DividendTimeLimit     int    `json:"dividendTimeLimit"`     // 至少几年分过红才参与排名
	AverageYear           int    `json:"averageYear"`           // 计算最近N年的平均分红收益率
	LastDividendYearLimit int    `json:"lastDividendYearLimit"` // 最近分红年份限制.如2018 表示至少2018分过红
}

type yieldItem struct {
	Year         int32
	Time         int64
	Yield        float64
	AveragePrice float64
	DividendCash float64
}

type yieldResult struct {
	Exchange         string       // 合约
	Symbol           string       // 交易所
	Name             string       // 股票名称
	UpdateTime       int64        // 更新时间
	DividendTimes    int32        // 分红次数，按年合并
	LastAverageYield float64      // 近三年平均收益
	LastYield        float64      // 最近一年分红收益率
	LastDividendYear int32        // 最后分红年份
	YieldList        []*yieldItem // 分红收益率
}

func getKlineArr(exchange, symbol string, pool *tdxclient.Pool) ([]*pb.Kline, error) {
	c, err := pool.GetQuoteClient()
	if err != nil {
		return nil, err
	}
	category := uint16(9)                             // 日线
	market := uint16(tdxclient.ToTdxMarket(exchange)) // 上海1， 深圳0

	var kArr []*pb.Kline
	var offset, page uint16
	page = 790
	for {
		a1, err := c.ReqGetSecurityBars(category, market, symbol, offset, page)
		if err != nil {
			return nil, err
		}
		if len(a1) < int(page) {
			break
		}
		l, err := tdxclient.ToKlineArr(a1)
		log.Printf("取K线[%d] %v %s %s", len(a1), err, time.Unix(l[0].Time, 0).Format("20060102"), time.Unix(l[len(l)-1].Time, 0).Format("20060102"))
		if err != nil {
			return nil, err
		}
		kArr = append(kArr, l...)
		offset += 790
	}
	return kArr, nil
}

func getKlienArr2(exchange, symbol string) ([]*pb.Kline, error) {
	var xq spider.Xueqiu
	return xq.KlineSeries(exchange, symbol, pb.PeriodType_D1, "normal", 0, 12000)
}

func yield(inst *pb.Instrument) (*yieldResult, error) {
	var res yieldResult
	res.Exchange = inst.Exchange
	res.Symbol = inst.Symbol
	res.Name = inst.Name
	res.UpdateTime = time.Now().Unix()
	log.Printf("[%s-%s]开始取K线", inst.Exchange, inst.Symbol)
	kArr, err := getKlienArr2(inst.Exchange, inst.Symbol)
	if err != nil {
		return nil, err
	}
	getAverage := func(start, end int64) float64 {
		// log.Printf("getaverage %s - %s", time.Unix(start, 0).Format("20060102"), time.Unix(end, 0).Format("20060102"))
		var res float64
		c := 0
		for _, k := range kArr {
			if k.Time >= start && k.Time < end {
				res += (k.Open + k.High + k.Low + k.Close) / 4
				c++
			}
		}
		if c > 0 {
			res = res / float64(c)
		}
		return res
	}
	var xq spider.Xueqiu
	historyBonus, err := xq.BonusHistory(inst.Exchange, inst.Symbol)
	if err != nil {
		log.Println(len(historyBonus), err)
		return nil, err
	}
	log.Printf("完成取K线和bonus,k线总长[%d] ", len(kArr))
	var lastDividendDate int64
	for _, b := range historyBonus {
		if b.DividendDate > 0 && b.DividendCash > 0 {
			if lastDividendDate > 0 {
				average := getAverage(lastDividendDate, b.DividendDate)
				if average > 0 {
					y := time.Unix(b.DividendDate, 0).Year()
					res.YieldList = append(res.YieldList, &yieldItem{Time: b.DividendDate, Yield: b.DividendCash / average, AveragePrice: average, DividendCash: b.DividendCash, Year: int32(y)})
				}
			}
			lastDividendDate = b.DividendDate
		}
	}
	var l []*yieldItem
	for _, y := range res.YieldList {
		f := false
		for _, py := range l {
			if y.Year == py.Year {
				// 同一年多次分红合并成一次
				py.Yield += y.Yield
				py.DividendCash += y.DividendCash
				f = true
				break
			}
		}
		if !f {
			l = append(l, y)
		}
	}
	res.YieldList = l
	return &res, nil
}

func yieldList() ([]*yieldResult, error) {
	var sse spider.SSE
	var sze spider.SZE
	var ret []*yieldResult
	l1, err := sse.StockList(false)
	if err != nil {
		return nil, err
	}
	l2, err := sze.StockList(false)
	if err != nil {
		return nil, err
	}
	l1 = append(l1, l2...)
	for _, inst := range l1 {
		y, err := yield(inst)
		if err == nil {
			ret = append(ret, y)
		} else {
			// log.Println(inst.Exchange, inst.Symbol, err)
		}
	}
	return ret, nil
}

func generateReport(c *bonusConfig, l []*yieldResult) error {
	path := c.OutputPath
	for _, r := range l {
		r.DividendTimes = int32(len(r.YieldList) + 1) // +1是因为第一次分红,还未计算其收益率，被屏蔽掉了
		if len(r.YieldList) > 0 {
			r.LastYield = r.YieldList[len(r.YieldList)-1].Yield
			r.LastDividendYear = r.YieldList[len(r.YieldList)-1].Year
		}
		if c.AverageYear > 0 && len(r.YieldList) > c.AverageYear {
			var aver float64
			var count int
			for i := (len(r.YieldList) - c.AverageYear); i >= 0; i-- {
				aver += r.YieldList[i].Yield
				count++
			}
			if count > 0 {
				aver /= float64(count)
				r.LastAverageYield = aver
			}
		}
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte{0xef, 0xbb, 0xbf})
	file.WriteString(fmt.Sprintf("交易所,代码,名称,最新股息率,最近%d年平均股息率,总分红年数,最新分红年份\r\n", c.AverageYear))
	for _, r := range l {
		if c.LastDividendYearLimit > 0 && int(r.LastDividendYear) < c.LastDividendYearLimit {
			continue
		}
		if int(r.DividendTimes) < c.DividendTimeLimit {
			continue
		}
		file.WriteString(fmt.Sprintf("%s,%s,%s,%.2f%%,%.2f%%,%d,%d\r\n", r.Exchange, r.Exchange+r.Symbol, r.Name, r.LastYield*100, r.LastAverageYield*100, r.DividendTimes, r.LastDividendYear))
	}
	return nil
}

func bonus() {
	var c bonusConfig
	var err error
	err = loadConfig("configs/bonus.json", &c)
	if err != nil {
		panic(err)
	}
	if c.CachePath == "" {
		panic("config.CachePath invalid")
	}
	if c.OutputPath == "" {
		panic("c.OutputPath invalid")
	}
	if c.DividendTimeLimit == 0 {
		c.DividendTimeLimit = 1
	}
	// 因为计算一次消耗时间长，所以缓存一个文件，如果需要生成最新报告，先删除 CachePath
	if _, err := os.Stat(c.CachePath); os.IsNotExist(err) {
		l, err := yieldList()
		if err != nil {
			panic(err)
		}
		f, err := os.Create(c.CachePath)
		if err != nil {
			log.Println("create file error ", err)
			panic(err)
		}
		out, err := json.Marshal(l)
		if err != nil {
			log.Println("Error:", err)
			panic(err)
		}
		_, err = f.Write(out)
		if err != nil {
			log.Println("Error:", err)
			panic(err)
		}
	}
	var l []*yieldResult
	file, err := os.Open(c.CachePath)
	if err != nil {
		log.Println("Error:", err)
		panic(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &l)
	if err != nil {
		panic(err)
	}
	// 输出报告
	generateReport(&c, l)
}

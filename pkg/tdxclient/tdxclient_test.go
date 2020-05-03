package tdxclient

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	pb "github.com/mineralres/protos/src/go/goshare"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func loadConfig(f string, out interface{}) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(data, &out)
}

func Test_i(t *testing.T) {
	var op PoolOptions
	err := loadConfig("config.json", &op)
	if err != nil {
		panic(err)
	}
	log.Println(op)
	pool := NewPool(&op)
	if pool == nil {
		panic(err)
	}
	client, err := pool.GetQuoteClient()
	if err != nil {
		panic(err)
	}
	// 查询股票数量
	count, err := client.ReqQryStockCount()
	if err != nil {
		panic(err)
	}
	log.Println(count, err)
	{
		// 查询股票列表
		arr, err := client.ReqGetSecurityList(0, 0)
		if err != nil {
			panic(err)
		}
		log.Println("查询股票列表", len(arr), err)
		for i := 0; i < len(arr) && i < 3; i++ {
			log.Println("股票:", arr[i])
		}
	}
	{
		// 查询行情
		arr, err := client.ReqGetSecurityQuotes([]*ReqGetInstrumentQuote{&ReqGetInstrumentQuote{Market: 0, Code: "000001"},
			&ReqGetInstrumentQuote{Market: 0, Code: "000002"}})
		if err != nil {
			panic(err)
		}
		if len(arr) != 2 {
			panic("should be 2")
		}
		tArr, err := toDepthsMarketDataArr(arr)
		if err != nil {
			panic(err)
		}
		log.Println("查询行情", len(arr), err)
		for _, m := range tArr {
			log.Println(m)
		}
	}
	{
		// 查询K线
		arr, err := client.ReqGetSecurityBars(9, 0, "000001", 0, 3)
		if err != nil {
			panic(err)
		}
		karr, err := ToKlineArr(arr)
		if err != nil {
			panic(err)
		}
		log.Println("查询K线", len(arr), err)
		for _, k := range karr {
			log.Println(k)
		}
	}
	{
		// 指数K线
		arr, err := client.ReqGetIndexBars(9, 1, "000001", 0, 3)
		if err != nil {
			panic(err)
		}
		karr, err := ToKlineArr(arr)
		if err != nil {
			panic(err)
		}
		log.Println("查询指数K线", len(arr), err)
		for _, k := range karr {
			log.Println(k)
		}
	}
	{
		// 分时数据
		arr, err := client.ReqGetMinuteTimeData(1, "600300")
		if err != nil {
			panic(err)
		}
		// for _, t := range arr {
		// 	log.Println(t)
		// }
		log.Println("查询分时行情", len(arr), err)
	}
	// 查询tick数据
}

func Test_quoteclient(t *testing.T) {
	return
	var op PoolOptions
	err := loadConfig("config.json", &op)
	if err != nil {
		panic(err)
	}
	log.Println(op)
	pool := NewPool(&op)
	if pool == nil {
		panic(err)
	}
	exclient, err := pool.GetExternClient()
	if err != nil {
		panic(err)
	}
	// 查市场
	marketList, err := exclient.GetMarketList()
	if err != nil {
		panic(err)
	}
	for i := range marketList {
		log.Println(marketList[i])
	}
	// 查询tick数据
	log.Println(exclient.GetLastTick("SHFE", "ru1909"))
	log.Println(exclient.GetLastTick("SHFE", "sc1909"))
	// 查询 50ETF期权 行情
	var req1 ReqGetInstrumentBars
	req1.Code = "10001886" // 8月购2850
	req1.Market = 8        // 上海个股期权
	req1.Category = ToTdxPeriod(pb.PeriodType_D1)
	req1.Start = uint32(0)
	req1.Count = uint16(100)
	ks, err := exclient.GetInstrumentBars(&req1)
	if err != nil {
		panic(err)
	}
	if len(ks) == 0 {
		panic("should not be 0")
	}
	log.Println("取长度", len(ks))
	// 查分时图行情
	ts, err := exclient.GetMinuteTimeData(8, "10001886")
	if err != nil {
		panic(err)
	}
	log.Println("取长度", len(ts))
}

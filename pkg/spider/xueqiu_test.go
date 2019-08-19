package spider

import (
	"log"
	"testing"
	"time"

	pb "github.com/mineralres/goshare/pkg/pb/goshare"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Test_i(t *testing.T) {
	var xq Xueqiu
	hl, err := xq.HotStockList()
	if err != nil {
		panic(err)
	}
	for _, item := range hl.GlobalH1 {
		log.Println(item)
	}
	starCount, err := xq.StarCount("SSE", "601318")
	if err != nil {
		panic(err)
	}
	log.Println("star count", starCount)
	{
		bonusList, err := xq.BonusHistory("SZE", "002008")
		if err != nil {
			panic(err)
		}
		for _, b := range bonusList {
			log.Println("分红配送", time.Unix(b.DividendDate, 0).Format("20060102"), b.PlanExplain)
		}
	}
	{
		// 取K线
		arr, err := xq.KlineSeries("SSE", "601398", pb.PeriodType_D1, "normal", 0, 9999)
		if err != nil {
			panic(err)
		}
		for i, k := range arr {
			log.Println(i, time.Unix(k.Time, 0).Format("20060102"), k.Open, k.High, k.Low, k.Close, k.Volume, k.Amount)
		}

	}
}

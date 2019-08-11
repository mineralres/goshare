package spider

import (
	"log"
	"testing"
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
}

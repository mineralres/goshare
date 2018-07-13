package main

import (
	"log"
	"time"

	"github.com/mineralres/goshare/pkg/goshare"
	"github.com/mineralres/goshare/pkg/pb"
)

func main() {
	// 新浪数据源
	var s goshare.SinaSource
	symbol := pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "rb1810"}
	// 获取5分钟历史K线
	data, err := s.GetKData(&symbol, pb.PeriodType_M5, 19990101, 20180307, 1)
	if err != nil {
		panic(err)
	}
	for i := range data.List {
		k := &data.List[i]
		log.Printf("%s %s [%.4f, %.4f, %.4f, %.4f]", time.Unix(k.Time, 0).Format("20060102 15:04:05"), symbol.Code, k.Open, k.High, k.Low, k.Close)
	}
	// 获取实时数据
	ret, err := s.GetLastTick(&symbol)
	log.Println(ret, err)
}

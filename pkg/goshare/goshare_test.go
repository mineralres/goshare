package goshare

import (
	"testing"

	"github.com/mineralres/goshare/aproto"
)

var s Service

func TestKData(t *testing.T) {
	symbol := aproto.Symbol{Exchange: aproto.ExchangeType_SHFE, Code: "rb1805"}
	k, err := s.GetKData(&symbol, aproto.PeriodType_M5, 19990101, 20180307, 1)
	if err != nil {
		t.Error(err)
	}
	if len(k.List) == 0 {
		t.Error("GetKData error")
	}
	// log.Println()
}

func TestGetLastTick(t *testing.T) {
	symbol := aproto.Symbol{Exchange: aproto.ExchangeType_SSE, Code: "600000"}
	md, err := s.GetLastTick(&symbol)
	if err != nil {
		t.Error(err)
	}
	if len(md.OrderBookList) == 0 {
		t.Error("获取行情盘口深度为空")
	}
	if md.Open == 0 {
		t.Error("md.Open == 0")
	}
	// log.Printf("Tick[%s], Open[%.2f], High[%.2f], Low[%.2f], Close[%.2f]", md.Symbol.Code, md.Open, md.High, md.Low, md.Close)
}

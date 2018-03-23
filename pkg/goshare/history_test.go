package goshare

import (
	"log"
	"testing"

	"github.com/mineralres/goshare/pkg/pb"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// TestGetKData TestGetKData
func TestGetKData(t *testing.T) {
	// symbol := pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "600000"}
	symbol := pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "rb1805"}
	var s Service
	ks, err := s.GetKData(&symbol, pb.PeriodType_M5, 19990101, 20180307, 1)
	if err != nil {
		t.Fatal(err)
	}
	if ks == nil {
		t.Fatal("查询K线失败")
	}
}

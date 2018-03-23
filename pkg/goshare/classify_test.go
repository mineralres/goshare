package goshare

import (
	"log"
	"testing"

	"github.com/mineralres/goshare/pkg/pb"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// TestIndexMemberData
func TestIndexMemberData(t *testing.T) {

	// symbol := pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "000016"}
	symbol := pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "000300"}
	var p Service
	arr_symbol, err := p.GetIndexMember(&symbol, 1)
	if err != nil {
		t.Error(err)
	}

	i := 0
	for i < len(arr_symbol) {
		log.Println(arr_symbol[i])
		i++
	}
	log.Println("000300 member size =", len(arr_symbol))
}

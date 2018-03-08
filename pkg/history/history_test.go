package history

import (
	"log"
	"testing"

	"github.com/mineralres/goshare/aproto"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// TestGetKData TestGetKData
func TestGetKData(t *testing.T) {
	symbol := aproto.Symbol{Exchange: aproto.ExchangeType_SSE, Code: "600000"}
	log.Println(GetKData(&symbol, aproto.PeriodType_D1, 19990101, 20180307, 1))
}

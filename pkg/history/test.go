package history

import (
	"log"

	"github.com/mineralres/goshare/aproto"
)

// TestGetKData TestGetKData
func TestGetKData() {
	symbol := aproto.Symbol{Exchange: aproto.ExchangeType_SSE, Code: "600000"}
	log.Println(GetKData(&symbol, aproto.PeriodType_D1, 19990101, 20180307, 1))
}

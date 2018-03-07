package realtime

import (
	"log"

	"github.com/mineralres/goshare/aproto"
)

// GetLastTick 取最新行情
func GetLastTick(symbol *aproto.Symbol) {
	log.Println("GetLastTick", symbol)
}

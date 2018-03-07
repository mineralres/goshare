package main

import (
	"log"

	"github.com/mineralres/goshare/aproto"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var symbol aproto.Symbol
	log.Println(symbol)
}

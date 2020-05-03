package main

import (
	"log"

	gspb "github.com/mineralres/protos/src/go/goshare"
)

func main() {
	var l gspb.Instrument
	l.Symbol = "ru2005"
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go serveWebsocket()
	db := &SixDB{}
	db.Run()
	serveHTTP(db)
}

func serveWebsocket() {
	port := 22345
	log.Printf("Websocket API on %d", port)
}

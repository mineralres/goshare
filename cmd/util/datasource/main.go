package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/mineralres/goshare/pkg/util/datasource"
)

func main() {
	go func() {
		// for debug http://localhost:6260/debug/pprof
		log.Println(http.ListenAndServe(":6260", nil))
	}()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	options := &datasource.Options{Token: "123465--", Port: 14999}
	front := datasource.MakeFront(options)
	front.Run()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
}

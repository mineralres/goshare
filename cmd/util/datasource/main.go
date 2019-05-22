package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mineralres/goshare/pkg/util/datasource"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	options := &datasource.Options{Token: "123465--", Port: 14999}
	front := datasource.MakeFront(options)
	front.Run()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
}

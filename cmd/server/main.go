package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go func() {
		// for debug http://localhost:6160/debug/pprof
		log.Println(http.ListenAndServe(":6160", nil))
	}()

	var c xconfig
	err := loadConfig("config.json", &c)
	if err != nil {
		panic(err)
	}
	go runGrpcService(c.GrpcPort)
	grpcEndPoint := fmt.Sprintf("localhost:%d", c.GrpcPort)
	go runGrpcGateway(c.GWHTTPPort, grpcEndPoint)
	wsf := makeWsFront(grpcEndPoint, c.WSPort)
	wsf.run()
}

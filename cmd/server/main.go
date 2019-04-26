package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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

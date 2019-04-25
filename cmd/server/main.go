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
	runGrpcGateway(c.GWHTTPPort, fmt.Sprintf("localhost:%d", c.GrpcPort))
}

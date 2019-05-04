package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os/exec"
	"time"
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
	go runGrpcService(c)
	go runGrpcGateway(c)
	go func() {
		http.Handle("/api/v1/", &proxyHandler{c: c})
		http.Handle("/", http.FileServer(http.Dir("./ui-release")))
		http.ListenAndServe(":9090", nil)
	}()
	go func() {
		time.Sleep(time.Second) // 等listen准备好
		cmd := exec.Command("explorer", "http://localhost:9090")
		err := cmd.Start()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
	wsf := makeWsFront(c)
	wsf.run()
}

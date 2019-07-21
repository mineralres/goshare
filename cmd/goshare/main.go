package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var c config
	err := loadConfig("config.json", &c)
	if err != nil {
		panic(err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go func() {
		// for debug http://localhost:6160/debug/pprof
		log.Println(http.ListenAndServe(":6160", nil))
	}()

	go func() {
		time.Sleep(time.Second) // 等listen准备好,打开默认浏览器
		cmd := exec.Command("explorer", "http://localhost:9090")
		err := cmd.Start()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	// 开启gin api router
	go func() {
		time.Sleep(time.Millisecond * 100)
		gw := NewGateway()
		gw.Run("./ui-release", 9090)
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	log.Println("所有服务退出, sig:", <-ch)
}

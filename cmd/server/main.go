package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/mineralres/goshare/pkg/api"
)

func main() {
	var c xconfig
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
	ctx, cancel := context.WithCancel(context.Background())

	// 开启服务
	go runSrv(ctx)
	// 开启gin api router
	go func() {
		r := api.MakeAPIRouter()
		http.Handle("/", r.Router)
		http.ListenAndServe(":9090", nil)
	}()
	// 退出信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	select {
	// wait on kill signal
	case <-ch:
		log.Println("this is kill signal")
		cancel()
	}
}

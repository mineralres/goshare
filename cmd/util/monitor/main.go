package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/mineralres/goshare/pkg/pb/goshare"
	"github.com/mineralres/goshare/pkg/spider"
	"github.com/mineralres/goshare/pkg/util"
)

// SMTPPassword 密码是自己在QQ邮箱的smtp授权码，参考 https://my.oschina.net/u/3768573/blog/1607327
type config struct {
	SMTPAccount  string   `json:"smtpAccount"`
	SMTPPassword string   `json:"smtpPassword"`
	SMTPServer   string   `json:"smtpServer"`
	Subscribers  []string `json:"subscribers"`
}

func loadConfig(f string, out interface{}) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(data, &out)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var c config
	err := loadConfig("config.json", &c)
	if err != nil {
		panic(err)
	}

	var spider spider.Spider
	symbol := &pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "601698"}
	for {
		tick, err := spider.GetLastTick(symbol)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(err, tick)
		if tick.Price <= 6.92 && false {
			for _, sub := range c.Subscribers {
				util.SendMail(c.SMTPAccount, c.SMTPPassword, c.SMTPServer, sub, "通知", "价格下跌")
			}
			break
		}
		time.Sleep(time.Second * 3)
	}
}

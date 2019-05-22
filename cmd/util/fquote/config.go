package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type config struct {
	BrokerID            string   `json:"brokerID"`
	Account             string   `json:"account"`
	Password            string   `json:"password"`
	TradeFrontList      []string `json:"tradeFrontList"`
	MarketDataFrontList []string `json:"marketDataFrontList"`
	Scheme              string   `json:"scheme"`
	Host                string   `json:"host"`
	Token               string   `json:"token"`
}

func loadConfig(f string, out interface{}) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(data, &out)
}

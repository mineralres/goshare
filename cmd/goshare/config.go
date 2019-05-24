package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/mineralres/goshare/pkg/service/dcenter/tdxclient"
)

type config struct {
	Port       int                   `json:"port"`
	TDXOptions tdxclient.PoolOptions `json:"tdxOptions"`
	GSURL      struct {
		Scheme string `json:"scheme"`
		Host   string `json:"host"`
		Token  string `json:"token"`
	} `json:"gsURL"`
}

func loadConfig(f string, out interface{}) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(data, &out)
}

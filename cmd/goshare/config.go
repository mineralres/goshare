package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type tdxOptions struct {
	ServerList       []string `json:"serverList"`
	ExternServerList []string `json:"externServerList"`
}

type config struct {
	Port       int        `json:"port"`
	TDXOptions tdxOptions `json:"tdxOptions"`
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

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type xservice struct {
	GrpcPort int `json:"grpcPort"`
	HTTPPort int `json:"httpPort"`
}

type xconfig struct {
	Common xservice `json:"common"`
	User   xservice `json:"user"`
	WSPort int      `json:"wsPort"`
}

func loadConfig(f string, out interface{}) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(data, &out)
}

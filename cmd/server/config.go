package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type xconfig struct {
	GrpcPort   int `json:"grpcPort"`
	GWHTTPPort int `json:"gwHTTPPort"`
	WSPort     int `json:"wsPort"`
}

func loadConfig(f string, out interface{}) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(data, &out)
}

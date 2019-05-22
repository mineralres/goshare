package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type config struct {
}

func loadConfig(f string, out interface{}) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(data, &out)
}

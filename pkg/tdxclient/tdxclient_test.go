package tdxclient

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func loadConfig(f string, out interface{}) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(data, &out)
}

func Test_i(t *testing.T) {
	var op PoolOptions
	err := loadConfig("config.json", &op)
	if err != nil {
		panic(err)
	}
	log.Println(op)
	pool := NewPool(&op)
	if pool == nil {
		panic(err)
	}
	client := pool.GetExternClient()
	log.Println(client.GetLastTick("SHFE", "ru1909"))
	log.Println(client.GetLastTick("SHFE", "sc1909"))
}

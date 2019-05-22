package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// PostSome post请求
func PostSome(url string, req, res interface{}) error {
	client := &http.Client{
		Timeout: time.Second * 3,
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(3 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*3)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	d, err := json.Marshal(req)
	if err != nil {
		return err
	}
	newreq, err := http.NewRequest("POST", url, bytes.NewReader(d))
	if err != nil {
		log.Println(err, url)
		return err
	}
	newreq.Close = true
	resp, err := client.Do(newreq)
	if err != nil {
		log.Println(err, url)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println(err, string(body))
		return err
	}
	return nil
}

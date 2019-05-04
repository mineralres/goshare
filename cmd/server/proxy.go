package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type proxyHandler struct {
	c xconfig
}

func (p *proxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error
	var req *http.Request
	client := &http.Client{}

	uri := fmt.Sprintf("http://localhost:%d", p.c.Common.HTTPPort) + r.RequestURI
	if strings.Index(r.RequestURI, "/api/v1/user") == 0 {
		uri = fmt.Sprintf("http://localhost:%d", p.c.User.HTTPPort) + r.RequestURI
	}
	// log.Println("转发请求", r.RequestURI, uri)
	req, err = http.NewRequest(r.Method, uri, r.Body)
	for name, value := range r.Header {
		req.Header.Set(name, value[0])
	}
	resp, err = client.Do(req)
	r.Body.Close()

	// combined for GET/POST
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

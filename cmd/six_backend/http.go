package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mineralres/goshare/pkg/util"
	"github.com/mineralres/protos/src/go/six"
)

func serveHTTP(db *SixDB) {
	port := 12345
	log.Printf("HTTP API on %d", port)
	h := func(path string, w http.ResponseWriter, req *http.Request) (interface{}, error) {
		switch path {
		case "/api/v1/tqv":
			return onTest(w, req, db)
		case "/api/v1/accountList":
			return onAccountList(w, req, db)
		}
		return nil, errors.New("API path not supported")
	}
	err := util.RunTinyGateway(port, h)
	if err != nil {
		panic(err)
	}
}

func onTest(w http.ResponseWriter, r *http.Request, db *SixDB) (interface{}, error) {
	return nil, nil
}

func onAccountList(w http.ResponseWriter, r *http.Request, db *SixDB) (interface{}, error) {
	return db.AccountList()
}

func onDeleteAccount(w http.ResponseWriter, r *http.Request, db *SixDB) (interface{}, error) {
	var req struct {
		BrokerID string `json:"brokerID"`
		UserID   string `json:"userID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("json err", err)
		return nil, err
	}
	return nil, db.DeleteAccount(req.BrokerID, req.UserID)
}

func onAddAccount(w http.ResponseWriter, r *http.Request, db *SixDB) (interface{}, error) {
	var req struct {
		BrokerID string `json:"brokerID"`
		UserID   string `json:"userID"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("json err", err)
		return nil, err
	}
	account := &six.Account{BrokerId: req.BrokerID, UserId: req.UserID, Password: req.Password}
	return nil, db.AddAccount(account)
}

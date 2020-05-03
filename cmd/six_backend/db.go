package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/mineralres/protos/src/go/six"
	"github.com/syndtr/goleveldb/leveldb"
)

// SixDB  six db CTP相关数据计算以及leveldb封装
type SixDB struct {
	db *leveldb.DB
}

// Run run
func (db *SixDB) Run() error {
	var err error
	if db.db, err = leveldb.OpenFile("sixdb", nil); err != nil {
		panic(err)
	}
	return err
}

// AccountList 交易账号列表
func (db *SixDB) AccountList() ([]*six.Account, error) {
	var ret []*six.Account
	ret = append(ret, &six.Account{UserId: "xiaobing", Password: "password"})
	return ret, nil
}

// AddAccount add account
func (db *SixDB) AddAccount(account *six.Account) error {
	return nil
}

// DeleteAccount delete account
func (db *SixDB) DeleteAccount(brokerID, userID string) error {
	return nil
}

// SetData set data
func (db *SixDB) SetData(key string, value proto.Message) error {
	data, err := proto.Marshal(value)
	if err != nil {
		log.Println(err)
		return err
	}
	return db.db.Put([]byte(key), data, nil)
}

// GetData get data
func (db *SixDB) GetData(key string, value proto.Message) error {
	data, err := db.db.Get([]byte(key), nil)
	if err != nil {
		return err
	}
	return proto.Unmarshal(data, value)
}

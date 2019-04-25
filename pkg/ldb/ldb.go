package ldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

// XLevelDB 存各种数据
type XLevelDB struct {
	// k线
	kdb *leveldb.DB
	// 按日存tick
	daytsdb *leveldb.DB
	// tick数据
	tickDB *leveldb.DB
}

// MakeXLevelDB Prepare
func MakeXLevelDB() *XLevelDB {
	var db XLevelDB
	var err error
	db.tickDB, err = leveldb.OpenFile("db/ldbtick", nil)
	if err != nil || db.tickDB == nil {
		log.Println(err)
		panic("open level db error")
	} else {
		log.Println("XLevelDB prepared")
	}
	db.kdb, err = leveldb.OpenFile("db/kdb", nil)
	if err != nil {
		panic("open kline leveldb error")
	}
	db.daytsdb, err = leveldb.OpenFile("db/daytsdb", nil)
	if err != nil {
		panic("open daytsdb leveldb error")
	}
	log.Println("KDB prepared")
	return &db
}

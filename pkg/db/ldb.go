package db

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/syndtr/goleveldb/leveldb"
)

// XLevelDB 存各种数据
type XLevelDB struct {
	// k线
	kdb *leveldb.DB
	// 按日存tick
	daytsdb *leveldb.DB
	// tick数据
	tickDB *leveldb.DB
	// 常规
	common *leveldb.DB
}

var (
	ErrSessionTimeout = errors.New("timeout")
)

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
	db.common, err = leveldb.OpenFile("db/common", nil)
	if err != nil {
		panic("open common leveldb error")
	}
	log.Println("KDB prepared")
	return &db
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// GetUniqueID 自增ID
func (db *XLevelDB) GetUniqueID() (int64, error) {
	var ret int64
	key := []byte("-x-unique-id")
	d, err := db.common.Get(key, nil)
	if err == nil {
		ret = BytesToInt64(d) + 1
	}
	d = Int64ToBytes(ret)
	err = db.common.Put(key, d, nil)
	return ret, err
}

// SetUserSession SetUserSession
func (db *XLevelDB) SetUserSession(token string, s *pb.UserSession) error {
	key := []byte(fmt.Sprintf("-token-%s", token))
	d, err := proto.Marshal(s)
	if err != nil {
		return err
	}
	return db.common.Put(key, d, nil)
}

// GetUserSession GetUserSession
func (db *XLevelDB) GetUserSession(token string) (*pb.UserSession, error) {
	key := []byte(fmt.Sprintf("-token-%s", token))
	d, err := db.common.Get(key, nil)
	if err != nil {
		return nil, err
	}
	var session pb.UserSession
	err = proto.Unmarshal(d, &session)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	if session.Deadline < now {
		return nil, ErrSessionTimeout
	}
	return &session, nil
}

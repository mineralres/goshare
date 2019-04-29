package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/gogo/protobuf/proto"
	"github.com/mineralres/goshare/pkg/pb"
)

// 保存行情
func (db *XLevelDB) saveTick(tick *pb.MarketDataSnapshot) error {
	key := fmt.Sprintf("%d-%s", tick.Symbol.Exchange, tick.Symbol.Code)
	out, _ := proto.Marshal(tick)
	err := db.tickDB.Put([]byte(key), out, nil)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 读取行情
func (db *XLevelDB) getTick(s *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	key := fmt.Sprintf("%d-%s", s.Exchange, s.Code)
	if db.tickDB == nil {
		log.Println("db.tickDB == nil")
		panic("db.tickDB == nil")
		return nil, errors.New("db.tickDB == nil")
	}
	data, err := db.tickDB.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}
	var ret pb.MarketDataSnapshot
	err = proto.Unmarshal(data, &ret)
	return &ret, err
}

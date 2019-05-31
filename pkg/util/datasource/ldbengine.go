package datasource

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"

	"github.com/mineralres/goshare/pkg/api"
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
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
	tradingInstrumentPrefix = "-tinstrument"
	tickPrefix              = "-tick"
)

// MakeXLevelDB Prepare
func MakeXLevelDB() *XLevelDB {
	var db XLevelDB
	var err error
	db.tickDB, err = leveldb.OpenFile("db/ldbtick", nil)
	if err != nil || db.tickDB == nil {
		log.Println(err)
		panic("open level db error")
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
	return &db
}

func int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func bytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// GetUniqueID 自增ID
func (db *XLevelDB) GetUniqueID() (int64, error) {
	var ret int64
	key := []byte("-x-unique-id")
	d, err := db.common.Get(key, nil)
	if err == nil {
		ret = bytesToInt64(d) + 1
	}
	d = int64ToBytes(ret)
	err = db.common.Put(key, d, nil)
	return ret, err
}

func makeKlineKey(s *pb.Symbol, period pb.PeriodType, t int64) string {
	return fmt.Sprintf("%d%s%d%d", s.Exchange, s.Code, int(period), t)
}

// Save 保存K线
func (db *XLevelDB) Save(s *pb.Symbol, period pb.PeriodType, k *pb.Kline) error {
	key := makeKlineKey(s, period, k.Time)
	d, err := proto.Marshal(k)
	if err != nil {
		return err
	}
	return db.kdb.Put([]byte(key), d, nil)
}

// SaveKlineSeries 存序列
func (db *XLevelDB) SaveKlineSeries(ks *pb.KlineSeries) error {
	for i := range ks.List {
		db.Save(ks.Symbol, ks.Period, ks.List[i])
	}
	return nil
}

func periodInSeconds(period pb.PeriodType) int32 {
	switch period {
	case pb.PeriodType_TICK:
		return 0
	case pb.PeriodType_M1:
		return 60
	case pb.PeriodType_M3:
		return 180
	case pb.PeriodType_M5:
		return 300
	case pb.PeriodType_M10:
		return 600
	case pb.PeriodType_M15:
		return 900
	case pb.PeriodType_M30:
		return 1800
	case pb.PeriodType_H1:
		return 3600
	case pb.PeriodType_H3:
		return 10800
	case pb.PeriodType_D1:
		return 86400
	case pb.PeriodType_W1:
		return 86400 * 7
	case pb.PeriodType_MON1:
		return 86400 * 30
	}
	return 0
}

// GetKlineSeries GetKlineSeries
func (db *XLevelDB) GetKlineSeries(ctx *api.Context, req *pb.ReqGetKlineSeries) (*pb.RspGetKlineSeries, error) {
	s := req.Symbol
	period := req.Period
	startTime := req.Start
	endTime := req.End
	lenLimit := req.LenLimit
	var ret pb.KlineSeries
	ret.Symbol = s
	ret.Period = period
	ret.PeriodInSeconds = periodInSeconds(period)
	keyStart := makeKlineKey(s, period, startTime)
	keyEnd := makeKlineKey(s, period, endTime)
	iter := db.kdb.NewIterator(&util.Range{Start: []byte(keyStart), Limit: []byte(keyEnd)}, nil)
	for iter.Next() {
		var k pb.Kline
		err := proto.Unmarshal(iter.Value(), &k)
		if err == nil {
			ret.List = append(ret.List, &k)
			if len(ret.List) >= int(lenLimit) {
				break
			}
		}
	}
	iter.Release()
	return &pb.RspGetKlineSeries{Symbol: s, Period: period, List: ret.List}, nil
}

// RGetKlineSeries 反向取
func (db *XLevelDB) RGetKlineSeries(ctx *api.Context, req *pb.ReqGetKlineSeries) (*pb.RspGetKlineSeries, error) {
	s := req.Symbol
	period := req.Period
	startTime := req.Start
	endTime := req.End
	lenLimit := req.LenLimit

	var ret pb.RspGetKlineSeries
	ret.Symbol = s
	ret.Period = period
	keyStart := makeKlineKey(s, period, startTime)
	keyEnd := makeKlineKey(s, period, endTime)
	iter := db.kdb.NewIterator(&util.Range{Start: []byte(keyStart), Limit: []byte(keyEnd)}, nil)
	for ok := iter.Last(); ok && (len(ret.List) < int(lenLimit)); ok = iter.Prev() {
		var k pb.Kline
		err := proto.Unmarshal(iter.Value(), &k)
		if err == nil {
			ret.List = append(ret.List, &k)
			if len(ret.List) >= int(lenLimit) {
				break
			}
		}
	}
	iter.Release()
	var rList []*pb.Kline
	for i := len(ret.List) - 1; i >= 0; i-- {
		rList = append(rList, ret.List[i])
	}
	ret.List = rList
	return &ret, nil
}

// SaveDayTickSeries 按日存ts
func (db *XLevelDB) SaveDayTickSeries(ts *pb.TickSeries) error {
	key := fmt.Sprintf("%d-%s-%d", ts.Symbol.Exchange, ts.Symbol.Code, ts.TradingDay)
	d, _ := proto.Marshal(ts)
	return db.daytsdb.Put([]byte(key), d, nil)
}

// 按日取ts
func (db *XLevelDB) getDayTickSeries(s *pb.Symbol, tradingDay int32) *pb.TickSeries {
	key := fmt.Sprintf("%d-%s-%d", s.Exchange, s.Code, tradingDay)
	d, err := db.daytsdb.Get([]byte(key), nil)
	if err != nil {
		return nil
	}
	var ret pb.TickSeries
	err = proto.Unmarshal(d, &ret)
	if err != nil {
		return nil
	}
	return &ret
}

// GetTickSerires 取最后一天的tick序列
func (db *XLevelDB) GetTickSerires(ctx *api.Context, req *pb.ReqGetTickSeries) (*pb.RspGetTickSeries, error) {
	s := req.Symbol
	keyStart := fmt.Sprintf("%d-%s-%d", s.Exchange, s.Code, 0)
	keyEnd := fmt.Sprintf("%d-%s-%d", s.Exchange, s.Code, 99999999)
	iter := db.daytsdb.NewIterator(&util.Range{Start: []byte(keyStart), Limit: []byte(keyEnd)}, nil)
	defer iter.Release()
	if iter.Last() {
		var ret pb.TickSeries
		err := proto.Unmarshal(iter.Value(), &ret)
		if err != nil {
			return nil, err
		}
		return &pb.RspGetTickSeries{List: ret.List}, nil
	}
	return &pb.RspGetTickSeries{}, nil
}

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
	}
	data, err := db.tickDB.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}
	var ret pb.MarketDataSnapshot
	err = proto.Unmarshal(data, &ret)
	return &ret, err
}

// SetTradingInstrument 保存合约信息
func (db *XLevelDB) SetTradingInstrument(inst *pb.TradingInstrument) error {
	key := fmt.Sprintf("%s-%d-%s", tradingInstrumentPrefix, inst.Symbol.Exchange, inst.Symbol.Code)
	d, err := proto.Marshal(inst)
	if err != nil {
		return err
	}
	return db.common.Put([]byte(key), d, nil)
}

// GetTradingInstrument 读取合约信息
func (db *XLevelDB) GetTradingInstrument(ctx *api.Context, s *pb.Symbol) (*pb.TradingInstrument, error) {
	var inst pb.TradingInstrument
	key := fmt.Sprintf("%s-%d-%s", tradingInstrumentPrefix, s.Exchange, s.Code)
	d, err := db.common.Get([]byte(key), nil)
	if err != nil {
		return &inst, err
	}
	err = proto.Unmarshal(d, &inst)
	return &inst, nil
}

// TradingInstrumentList 全部合约信息
func (db *XLevelDB) TradingInstrumentList(ctx *api.Context, req *pb.ReqGetTradingInstrumentList) ([]*pb.TradingInstrument, error) {
	log.Print("TradingInstrumentList ldb")
	var ret []*pb.TradingInstrument
	iter := db.common.NewIterator(util.BytesPrefix([]byte(tradingInstrumentPrefix)), nil)
	defer iter.Release()
	for iter.Next() {
		var inst pb.TradingInstrument
		err := proto.Unmarshal(iter.Value(), &inst)
		if err == nil {
			ret = append(ret, &inst)
		} else {
			log.Println(err)
		}
	}
	if len(ret) == 0 {
		return ret, errors.New("empty")
	}
	return ret, nil
}

// SetTick SetTick
func (db *XLevelDB) SetTick(req *pb.MarketDataSnapshot) error {
	key := fmt.Sprintf("%s-%d-%s", tickPrefix, req.Symbol.Exchange, req.Symbol.Code)
	d, err := proto.Marshal(req)
	if err != nil {
		return err
	}
	return db.common.Put([]byte(key), d, nil)
}

// GetLastTick GetTick
func (db *XLevelDB) GetLastTick(ctx *api.Context, req *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	var tick pb.MarketDataSnapshot
	key := fmt.Sprintf("%s-%d-%s", tickPrefix, req.Exchange, req.Code)
	d, err := db.common.Get([]byte(key), nil)
	if err != nil {
		return &tick, err
	}
	err = proto.Unmarshal(d, &tick)
	return &tick, nil
}

// SetMainContract 保存合约信息
func (db *XLevelDB) SetMainContract(day int32, l *pb.TradingInstrumentList) error {
	key := fmt.Sprintf("-main-contract-%d", day)
	d, err := proto.Marshal(l)
	if err != nil {
		return err
	}
	return db.common.Put([]byte(key), d, nil)
}

// GetMainContract 读取合约信息
func (db *XLevelDB) GetMainContract(ctx *api.Context, day int32) (*pb.TradingInstrumentList, error) {
	var l pb.TradingInstrumentList
	key := fmt.Sprintf("-main-contract-%d", day)
	d, err := db.common.Get([]byte(key), nil)
	if err != nil {
		return &l, err
	}
	err = proto.Unmarshal(d, &l)
	return &l, nil
}
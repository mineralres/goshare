package util

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"

	pb "github.com/mineralres/goshare/pkg/pb/goshare"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// LdbCache 存各种数据
type LdbCache struct {
	kdb     *leveldb.DB // k线
	daytsdb *leveldb.DB // 按日存tick
	tickDB  *leveldb.DB // tick数据
	tiDB    *leveldb.DB // 合约
}

// NewLdbCache Prepare
func NewLdbCache() (*LdbCache, error) {
	var db LdbCache
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
	db.tiDB, err = leveldb.OpenFile("db/tidb", nil)
	if err != nil {
		panic("open tidb leveldb error")
	}
	return &db, err
}

func int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func bytesToInt64(buf []byte) (int64, error) {
	if len(buf) != 8 {
		return 0, errors.New("input array should be 8 bytes")
	}
	return int64(binary.BigEndian.Uint64(buf)), nil
}

func makeKlineKey(ex, symbol string, period pb.PeriodType, t int64) string {
	return fmt.Sprintf("%s%s%d%d", ex, symbol, int(period), t)
}

// Save 保存K线
func (db *LdbCache) Save(ex, symbol string, period pb.PeriodType, k *pb.Kline) error {
	key := makeKlineKey(ex, symbol, period, k.Time)
	d, err := proto.Marshal(k)
	if err != nil {
		return err
	}
	return db.kdb.Put([]byte(key), d, nil)
}

// SaveKlineSeries 存序列
func (db *LdbCache) SaveKlineSeries(ks *pb.KlineSeries) error {
	for i := range ks.List {
		db.Save(ks.Exchange, ks.Symbol, ks.Period, ks.List[i])
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
func (db *LdbCache) GetKlineSeries(req *pb.ReqGetKlineSeries) (*pb.RspGetKlineSeries, error) {
	ex := req.Exchange
	symbol := req.Symbol
	period := req.Period
	startTime := req.Start
	endTime := req.End
	lenLimit := req.LenLimit
	var ret pb.KlineSeries
	ret.Exchange = ex
	ret.Symbol = symbol
	ret.Period = period
	ret.PeriodInSeconds = periodInSeconds(period)
	keyStart := makeKlineKey(ex, symbol, period, startTime)
	keyEnd := makeKlineKey(ex, symbol, period, endTime)
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
	return &pb.RspGetKlineSeries{Exchange: ex, Symbol: symbol, Period: period, List: ret.List}, nil
}

// RGetKlineSeries 反向取
func (db *LdbCache) RGetKlineSeries(req *pb.ReqGetKlineSeries) (*pb.RspGetKlineSeries, error) {
	ex := req.Exchange
	symbol := req.Symbol
	period := req.Period
	startTime := req.Start
	endTime := req.End
	lenLimit := req.LenLimit

	var ret pb.RspGetKlineSeries
	ret.Exchange = ex
	ret.Symbol = symbol
	ret.Period = period
	keyStart := makeKlineKey(ex, symbol, period, startTime)
	keyEnd := makeKlineKey(ex, symbol, period, endTime)
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

// SetDayTickSeries 按日存ts
func (db *LdbCache) SetDayTickSeries(ts *pb.TickSeries) error {
	key := fmt.Sprintf("%s-%s-%d", ts.Exchange, ts.Symbol, ts.TradingDay)
	d, _ := proto.Marshal(ts)
	return db.daytsdb.Put([]byte(key), d, nil)
}

// GetDayTickSeries 按日取ts
func (db *LdbCache) GetDayTickSeries(ex, symbol string, tradingDay int32) *pb.TickSeries {
	key := fmt.Sprintf("%s-%s-%d", ex, symbol, tradingDay)
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

// GetLatestTickSerires 取最后一天的tick序列
func (db *LdbCache) GetLatestTickSerires(req *pb.ReqGetTickSeries) (*pb.RspGetTickSeries, error) {
	keyStart := fmt.Sprintf("%s-%s-%d", req.Exchange, req.Symbol, 0)
	keyEnd := fmt.Sprintf("%s-%s-%d", req.Exchange, req.Symbol, 99999999)
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

// SetTick 保存行情
func (db *LdbCache) SetTick(tick *pb.MarketDataSnapshot) error {
	key := fmt.Sprintf("%s-%s", tick.Exchange, tick.Symbol)
	out, _ := proto.Marshal(tick)
	err := db.tickDB.Put([]byte(key), out, nil)
	if err != nil {
		log.Println(err)
	}
	return err
}

// GetTick 读取行情
func (db *LdbCache) GetTick(ex, symbol string) (*pb.MarketDataSnapshot, error) {
	key := fmt.Sprintf("%s-%s", ex, symbol)
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

// SetInstrument 保存合约信息
func (db *LdbCache) SetInstrument(inst *pb.Instrument) error {
	if inst.Exchange == "" && inst.Symbol == "" {
		return errors.New("empty symbol")
	}
	key := fmt.Sprintf("%s-%s", inst.Exchange, inst.Symbol)
	d, err := proto.Marshal(inst)
	if err != nil {
		return err
	}
	return db.tiDB.Put([]byte(key), d, nil)
}

// GetInstrument 读取合约信息
func (db *LdbCache) GetInstrument(ex, symbol string) (*pb.Instrument, error) {
	var inst pb.Instrument
	key := fmt.Sprintf("%s-%s", ex, symbol)
	d, err := db.tiDB.Get([]byte(key), nil)
	if err != nil {
		return &inst, err
	}
	err = proto.Unmarshal(d, &inst)
	return &inst, nil
}

// InstrumentList 全部合约信息
func (db *LdbCache) InstrumentList(req *pb.ReqGetTradingInstrumentList) ([]*pb.Instrument, error) {
	log.Print("InstrumentList ldb")
	var ret []*pb.Instrument
	iter := db.tiDB.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		var inst pb.Instrument
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

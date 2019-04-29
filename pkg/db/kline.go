package db

import (
	"fmt"

	"github.com/mineralres/goshare/pkg/pb"

	"github.com/gogo/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb/util"
)

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

func PeriodInSeconds(period pb.PeriodType) int32 {
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
func (db *XLevelDB) GetKlineSeries(s *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) *pb.KlineSeries {
	var ret pb.KlineSeries
	ret.Symbol = s
	ret.Period = period
	ret.PeriodInSeconds = PeriodInSeconds(period)
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
	return &ret
}

// RGetKlineSeries 反向取
func (db *XLevelDB) RGetKlineSeries(s *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) *pb.KlineSeries {
	var ret pb.KlineSeries
	ret.Symbol = s
	ret.Period = period
	ret.PeriodInSeconds = PeriodInSeconds(period)
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
	return &ret
}

// 按日存ts
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

// GetLastTickSerires 取最后一天的tick序列
func (db *XLevelDB) GetLastTickSerires(s *pb.Symbol) *pb.TickSeries {
	keyStart := fmt.Sprintf("%d-%s-%d", s.Exchange, s.Code, 0)
	keyEnd := fmt.Sprintf("%d-%s-%d", s.Exchange, s.Code, 99999999)
	iter := db.daytsdb.NewIterator(&util.Range{Start: []byte(keyStart), Limit: []byte(keyEnd)}, nil)
	defer iter.Release()
	if iter.Last() {
		var ret pb.TickSeries
		err := proto.Unmarshal(iter.Value(), &ret)
		if err != nil {
			return nil
		}
		return &ret
	}
	return nil
}

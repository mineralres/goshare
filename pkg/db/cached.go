package db

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/mineralres/goshare/pkg/pb"
)

// XCache 内存中K线
type XCache struct {
	xdb         XDataBase
	ksMap       sync.Map
	chLastKline chan lastKline
	chLastDayTS chan pb.TickSeries
}

// XDB 底层数据库
func (cache *XCache) XDB() XDataBase {
	return cache.xdb
}

// 最新一根K线，用于保存到leveldb
type lastKline struct {
	s      pb.Symbol
	period pb.PeriodType
	kline  pb.Kline
}

// cachedSymbolData cachedSymbolData
type cachedSymbolData struct {
	sync.RWMutex
	Symbol      pb.Symbol
	KlineList   []pb.KlineSeries
	chLastKline chan lastKline
	chLastDayTS chan pb.TickSeries
	db          XDataBase
	cacheLen    int64
	dayTS       pb.TickSeries
	tick        *pb.MarketDataSnapshot
}

func (sd *cachedSymbolData) init(s pb.Symbol, ch chan lastKline, db XDataBase, chts chan pb.TickSeries) {
	sd.Symbol = s
	sd.chLastKline = ch
	sd.chLastDayTS = chts
	sd.db = db
	sd.cacheLen = 1000
	sd.tick = &pb.MarketDataSnapshot{}
	sd.tick.Symbol = &s
	nowT := time.Now().Unix()
	for i := 1; i < 100; i++ {
		if i > int(pb.PeriodType_MON1) {
			break
		}
		ks := db.RGetKlineSeries(&s, pb.PeriodType(i), 0, nowT, sd.cacheLen)
		if ks == nil {
			ks = &pb.KlineSeries{Period: pb.PeriodType(i), Symbol: &s}
		} else {
			// log.Println("db.RGetKlineSeries", s.Code, pb.PeriodType(i), len(ks.List))
		}
		ks.PeriodInSeconds = PeriodInSeconds(ks.Period)
		if ks.PeriodInSeconds == 0 {
			panic("ks.PeriodInSeconds == 0")
		}
		sd.KlineList = append(sd.KlineList, *ks)
	}
	ts := sd.db.GetLastTickSerires(&s)
	if ts != nil {
		sd.dayTS = *ts
		// log.Println("getLastTickSerires", s.Code, len(sd.dayTS.List))
	}
}

// MakeXCache MakeXCache
func MakeXCache() *XCache {
	var c XCache
	c.xdb = MakeXLevelDB()
	c.chLastKline = make(chan lastKline, 19999)
	c.chLastDayTS = make(chan pb.TickSeries, 1000)
	// 保存K线
	go c.saver()
	// 每分钟保存
	go c.flushChecker()

	return &c
}

func (cache *XCache) saver() {
	for {
		select {
		case k := <-cache.chLastKline:
			cache.xdb.Save(&k.s, k.period, &k.kline)
			break
		case ts := <-cache.chLastDayTS:
			cache.xdb.SaveDayTickSeries(&ts)
			break
		}
	}
}

func (cache *XCache) flushChecker() {
	for {
		time.Sleep(time.Second * 60)
		cache.ksMap.Range(func(k, v interface{}) bool {
			sd := v.(*cachedSymbolData)
			sd.flush()
			return true
		})
	}
}

func (sd *cachedSymbolData) flush() {
	sd.RLock()
	for i := range sd.KlineList {
		ks := &sd.KlineList[i]
		if len(ks.List) > 0 {
			sd.chLastKline <- lastKline{s: sd.Symbol, period: ks.Period, kline: *ks.List[len(ks.List)-1]}
		}
	}
	ts := sd.dayTS
	ts.List = make([]*pb.MarketDataSnapshot, len(sd.dayTS.List))
	ts.Symbol = &sd.Symbol
	copy(ts.List, sd.dayTS.List)
	sd.chLastDayTS <- ts
	sd.RUnlock()
}

func (sd *cachedSymbolData) updateKlineSeries(tick *pb.MarketDataSnapshot) {
	sd.Lock()
	defer sd.Unlock()
	sd.tick = tick
	if len(sd.dayTS.List) == 0 || tick.TradingDay > sd.dayTS.TradingDay {
		// 重新初始化
		sd.dayTS = pb.TickSeries{}
		sd.dayTS.TradingDay = tick.TradingDay
		sd.dayTS.Symbol = tick.Symbol
		sd.dayTS.List = append(sd.dayTS.List, tick)
		tick.VolumeDelta = tick.Volume
	} else if tick.TradingDay < sd.dayTS.TradingDay {
		return
	}
	if len(sd.dayTS.List) > 1 {
		pre := sd.dayTS.List[len(sd.dayTS.List)-2]
		if tick.Time < pre.Time {
			// 时间不能小
			return
		} else if tick.Time == pre.Time {
			if tick.Price == pre.Price && tick.Volume == pre.Volume {
				// 没有变化
				// log.Println("ret2", tick.Symbol.Code, tick.TradingDay, sd.dayTS.TradingDay)
				// return
			}
		}
		tick.VolumeDelta = tick.Volume - pre.Volume
	}
	sd.dayTS.List = append(sd.dayTS.List, tick)
	if len(sd.dayTS.List) > 1 {
		sd.dayTS.List = sd.dayTS.List[len(sd.dayTS.List)-2:]
	}

	for i := range sd.KlineList {
		ks := &sd.KlineList[i]
		preLen := len(ks.List)
		UpdateKlineSeries(ks, tick)
		if len(ks.List) > preLen {
			// log.Println(tick.Symbol.Code, ks.Period, len(ks.List), "sd.chLastKline")
			sd.chLastKline <- lastKline{s: *tick.Symbol, period: ks.Period, kline: *ks.List[preLen]}
		}
	}
}

func (sd *cachedSymbolData) getLastTick() *pb.MarketDataSnapshot {
	sd.RLock()
	defer sd.RUnlock()
	return sd.tick
}

func (sd *cachedSymbolData) getKlineSeries(period pb.PeriodType, startTime, endTime, lenLimit int64) *pb.KlineSeries {
	sd.RLock()
	defer sd.RUnlock()
	ret := &pb.KlineSeries{}
	ret.Symbol = &sd.Symbol
	ret.Period = period
	// 如果长度大于cacheLen就是leveldb中读取
	for i := range sd.KlineList {
		if sd.KlineList[i].Period == period {
			// ret.List = make([]pb.Kline, len(sd.KlineList[i].List))
			l := sd.KlineList[i].List
			for j := range l {
				k := l[j]
				if validKline(k) {
					ret.List = append(ret.List, k)
				}
			}
			// copy(ret.List, sd.KlineList[i].List)
			// log.Println("getKlineSeries", period, startTime, endTime, lenLimit, len(sd.KlineList[i].List), len(ret.List))
		}
	}
	return ret
}

func (sd *cachedSymbolData) getLastKline(period pb.PeriodType) pb.Kline {
	for i := range sd.KlineList {
		ks := &sd.KlineList[i]
		if ks.Period == period {
			if len(ks.List) > 0 {
				return *ks.List[len(ks.List)-1]
			}
		}
	}
	return pb.Kline{}
}

// Update UpdateKlineSeries
func (cache *XCache) Update(tick *pb.MarketDataSnapshot) {
	if tick.Price <= 0 || tick.Price > 9999999 {
		return
	}
	var sd *cachedSymbolData
	v, ok := cache.ksMap.Load(tick.Symbol)
	if ok {
		sd = v.(*cachedSymbolData)
	} else {
		// 添加新的group
		sd = &cachedSymbolData{}
		sd.init(*tick.Symbol, cache.chLastKline, cache.xdb, cache.chLastDayTS)
		cache.ksMap.Store(tick.Symbol, sd)
	}
	sd.updateKlineSeries(tick)
}

// GetKlineSeries 取K线
func (cache *XCache) GetKlineSeries(symbol *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) *pb.KlineSeries {
	if symbol.Code == "" {
		return nil
	}
	// log.Println("getKlineSeries", symbol, period, startTime, endTime, lenLimit)
	v, ok := cache.ksMap.Load(*symbol)
	var sd *cachedSymbolData
	if !ok {
		// 再从leveldb里找
		// 添加新的group
		sd = &cachedSymbolData{}
		sd.init(*symbol, cache.chLastKline, cache.xdb, cache.chLastDayTS)
		cache.ksMap.Store(*symbol, sd)
	} else {
		sd = v.(*cachedSymbolData)
	}
	return sd.getKlineSeries(period, startTime, endTime, lenLimit)
}

// GetLastTick 最新盘口
func (cache *XCache) GetLastTick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	if symbol.Code == "" {
		return nil, errors.New("empty code")
	}
	// log.Println("getKlineSeries", symbol, period, startTime, endTime, lenLimit)
	v, ok := cache.ksMap.Load(*symbol)
	if ok {
		return v.(*cachedSymbolData).getLastTick(), nil
	}
	return nil, errors.New("no cache data")
}

// GetLastKline GetLastKline最新K线
func (cache *XCache) GetLastKline(symbol *pb.Symbol, period pb.PeriodType) *pb.Kline {
	if symbol.Code == "" {
		return nil
	}
	// log.Println("getKlineSeries", symbol, period, startTime, endTime, lenLimit)
	v, ok := cache.ksMap.Load(*symbol)
	var sd *cachedSymbolData
	if !ok {
		return nil
	}
	sd = v.(*cachedSymbolData)
	k := sd.getLastKline(period)
	if k.Time > 0 {
		return &k
	}
	return nil
}

func makeKline(tick *pb.MarketDataSnapshot, baseTime int64) *pb.Kline {
	var kline pb.Kline
	kline.Open = tick.Price
	kline.High = tick.Price
	kline.Low = tick.Price
	kline.Close = tick.Price
	kline.Time = baseTime
	kline.Volume = tick.VolumeDelta
	kline.Position = tick.Position
	kline.TradingDay = tick.TradingDay
	return &kline
}

// UpdateKlineSeries UpdateKlineSeries
func UpdateKlineSeries(ks *pb.KlineSeries, tick *pb.MarketDataSnapshot) {
	if tick.Price <= 0 || tick.Price > 9999999 {
		return
	}
	if ks.PeriodInSeconds <= 0 {
		log.Println("ks.PeriodInSeconds <= 0", ks.Period, ks.PeriodInSeconds, tick.Symbol.Code)
		panic("ks.PeriodInSeconds <= 0")
	}
	baseTime := int64(tick.Time/int64(ks.PeriodInSeconds)) * int64(ks.PeriodInSeconds)
	if len(ks.List) == 0 {
		ks.List = append(ks.List, makeKline(tick, baseTime))
		return
	}
	lastKline := ks.List[len(ks.List)-1]
	if baseTime > lastKline.Time {
		ks.List = append(ks.List, makeKline(tick, baseTime))
		return
	}
	if tick.Price > lastKline.High {
		lastKline.High = tick.Price
	}
	if tick.Price < lastKline.Low {
		lastKline.Low = tick.Price
	}
	lastKline.Close = tick.Price
	lastKline.Position = tick.Position
	lastKline.Volume += tick.VolumeDelta
}

// GetSimpleTickSeries GetSimpleTickSeries
func (cache *XCache) GetSimpleTickSeries(s *pb.Symbol) *pb.SimpleTickSeries {
	ks := cache.GetKlineSeries(s, pb.PeriodType_M1, 0, time.Now().Unix(), 999)
	if ks == nil {
		return nil
	}
	if len(ks.List) == 0 {
		return nil
	}
	tradingDay := ks.List[len(ks.List)-1].TradingDay
	var ret pb.SimpleTickSeries
	for i := range ks.List {
		k := ks.List[i]
		if k.TradingDay == tradingDay && validKline(k) {
			ret.List = append(ret.List, &pb.SimpleTick{Price: k.Close, Time: k.Time, Volume: k.Volume})
		}
	}
	return &ret
}

func validKline(k *pb.Kline) bool {
	pmax := 99999999.99
	if k.Time == 0 || k.Volume == 0 {
		return false
	}
	if k.Open == 0 && k.High == 0 && k.Low == 0 && k.Close == 0 {
		return false
	}
	if k.Open > pmax || k.Open < 0 {
		return false
	}
	if k.High > pmax || k.High < 0 {
		return false
	}
	if k.Low > pmax || k.Low < 0 {
		return false
	}
	if k.Close > pmax || k.Close < 0 {
		return false
	}
	return true
}

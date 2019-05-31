package datasource

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/mineralres/goshare/pkg/api"
	"github.com/mineralres/goshare/pkg/pb"
)

// Backend 数据存储引擎
type Backend interface {
	// 保存单根K线
	Save(*pb.Symbol, pb.PeriodType, *pb.Kline) error
	// 保存日tick序列
	SaveDayTickSeries(ts *pb.TickSeries) error
	// 保存合约信息
	SetTradingInstrument(*pb.TradingInstrument) error
	// 保存tick
	SetTick(*pb.MarketDataSnapshot) error
	// 保存主力合约
	SetMainContract(day int32, l *pb.TradingInstrumentList) error
	// 读取主力合约
	GetMainContract(ctx *api.Context, day int32) (*pb.TradingInstrumentList, error)
}

func makeKey(s *pb.Symbol) string {
	return fmt.Sprintf("%d-%s", s.Exchange, s.Code)
}

// XCache 内存中K线
type XCache struct {
	ds        api.DataSource         // 数据源
	backend   Backend                // 存储
	sdMap     map[string]*symbolData // sdmap
	sdMapLock sync.RWMutex           // 锁
}

// 最新一根K线，用于保存到leveldb
type lastKline struct {
	s      pb.Symbol
	period pb.PeriodType
	kline  pb.Kline
}

// symbolData symbolData
type symbolData struct {
	sync.RWMutex
	Symbol      pb.Symbol                     // 合约
	KlineList   []pb.KlineSeries              // K线
	dayTS       pb.TickSeries                 // tick
	tick        *pb.MarketDataSnapshot        // 盘口
	cache       *XCache                       //
	subscribers []chan *pb.MarketDataSnapshot // 订阅
	ti          pb.TradingInstrument          // 属性
}

func (sd *symbolData) subscribe(ch chan *pb.MarketDataSnapshot) {
	sd.Lock()
	defer sd.Unlock()
	for i := range sd.subscribers {
		if ch == sd.subscribers[i] {
			// 重复订阅
			return
		}
	}
	sd.subscribers = append(sd.subscribers, ch)
	ch <- sd.tick
}

func (sd *symbolData) unsubscribe(ch chan *pb.MarketDataSnapshot) {
	sd.Lock()
	defer sd.Unlock()
	var left []chan *pb.MarketDataSnapshot
	for i := range sd.subscribers {
		if ch != sd.subscribers[i] {
			left = append(left, sd.subscribers[i])
		}
	}
	sd.subscribers = left
}

func initSymbolData(s pb.Symbol, cache *XCache) *symbolData {
	sd := &symbolData{}
	ti, err := cache.getTradingInstrument(&pb.ReqGetTradingInstrument{Symbol: &s})
	if err != nil {
		return nil
	}
	sd.ti = *ti
	sd.Symbol = s
	sd.cache = cache
	tick, err := sd.cache.getTick(&s)
	if err == nil {
		sd.tick = tick
	} else {
		sd.tick = &pb.MarketDataSnapshot{}
	}
	sd.tick.Symbol = &s
	nowT := time.Now().Unix()
	for i := 1; i < 100; i++ {
		if i > int(pb.PeriodType_MON1) {
			break
		}
		ks, err := cache.rGetKlineSeries(&s, pb.PeriodType(i), 0, nowT, 2)
		if err != nil {
			ks = &pb.KlineSeries{Period: pb.PeriodType(i), Symbol: &s}
		} else {
			// log.Println("db.RGetKlineSeries", s.Code, pb.PeriodType(i), len(ks.List))
		}
		ks.PeriodInSeconds = periodInSeconds(ks.Period)
		if ks.PeriodInSeconds == 0 {
			log.Println("ks.PeriodInSeconds", ks.PeriodInSeconds, ks.Period, pb.PeriodType(i))
			panic("ks.PeriodInSeconds == 0")
		}
		sd.KlineList = append(sd.KlineList, *ks)
	}
	ts, err := cache.getLastTickSerires(&s)
	if err == nil {
		sd.dayTS = *ts
	}
	return sd
}

func makeXCache(ds api.DataSource, backend Backend) *XCache {
	var c XCache
	c.sdMap = make(map[string]*symbolData)
	c.ds = ds
	c.backend = backend
	// 每分钟保存
	go c.flushChecker()
	return &c
}

func (cache *XCache) flushChecker() {
	// 定时把最新k线和tick放进保存队列
	for {
		time.Sleep(time.Second * 60)
		var l []*symbolData
		cache.sdMapLock.RLock()
		for _, sd := range cache.sdMap {
			l = append(l, sd)
		}
		cache.sdMapLock.RUnlock()
		for i := range l {
			l[i].flush()
		}
	}
}

func (sd *symbolData) flush() {
	sd.Lock()
	// 保存K线
	for i := range sd.KlineList {
		ks := &sd.KlineList[i]
		sz := len(ks.List)
		for i := range ks.List {
			sd.cache.backend.Save(&sd.Symbol, ks.Period, ks.List[i])
		}
		if sz > 1 {
			ks.List = ks.List[sz-1:]
		}
	}
	// daily ts
	sd.Unlock()
}

func (sd *symbolData) update(tick *pb.MarketDataSnapshot) {
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
			if tick.Price == pre.Price && tick.Volume == pre.Volume && tick.Position == pre.Position {
				// 没有变化
				// log.Println("ret2", tick.Symbol.Code, tick.TradingDay, sd.dayTS.TradingDay)
				return
			}
		}
		tick.VolumeDelta = tick.Volume - pre.Volume
	}
	// sd.dayTS.List = append(sd.dayTS.List, tick)
	for i := range sd.KlineList {
		ks := &sd.KlineList[i]
		UpdateKlineSeries(ks, tick)
	}
	for i := range sd.subscribers {
		sd.subscribers[i] <- tick
	}
	sd.cache.backend.SetTick(tick)
}

func (sd *symbolData) getLastTick() *pb.MarketDataSnapshot {
	sd.RLock()
	defer sd.RUnlock()
	return sd.tick
}

func (sd *symbolData) getKlineSeries(period pb.PeriodType, startTime, endTime, lenLimit int64) *pb.KlineSeries {
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

func (sd *symbolData) getLastKline(period pb.PeriodType) pb.Kline {
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

func (cache *XCache) getSymbolData(s *pb.Symbol) *symbolData {
	cache.sdMapLock.Lock()
	v, ok := cache.sdMap[makeKey(s)]
	if !ok {
		v = initSymbolData(*s, cache)
		if v != nil {
			cache.sdMap[makeKey(s)] = v
		}
	}
	cache.sdMapLock.Unlock()
	return v
}

// Update UpdateKlineSeries
func (cache *XCache) Update(tick *pb.MarketDataSnapshot) {
	if tick.Price <= 0 || tick.Price > 9999999 {
		return
	}
	sd := cache.getSymbolData(tick.Symbol)
	if sd == nil {
		return
	}
	sd.update(tick)
}

// GetKlineSeries 取K线
func (cache *XCache) GetKlineSeries(symbol *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) *pb.KlineSeries {
	if symbol.Code == "" {
		return nil
	}
	// log.Println("getKlineSeries", symbol, period, startTime, endTime, lenLimit)
	sd := cache.getSymbolData(symbol)
	return sd.getKlineSeries(period, startTime, endTime, lenLimit)
}

// GetLastTick 最新盘口
func (cache *XCache) GetLastTick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	if symbol.Code == "" {
		return nil, errors.New("empty code")
	}
	// log.Println("getKlineSeries", symbol, period, startTime, endTime, lenLimit)
	cache.sdMapLock.Lock()
	defer cache.sdMapLock.Unlock()
	v, ok := cache.sdMap[makeKey(symbol)]
	if ok {
		return v.getLastTick(), nil
	}
	return nil, errors.New("no cache data")
}

// GetLastKline GetLastKline最新K线
func (cache *XCache) GetLastKline(symbol *pb.Symbol, period pb.PeriodType) *pb.Kline {
	if symbol.Code == "" {
		return nil
	}
	// log.Println("getKlineSeries", symbol, period, startTime, endTime, lenLimit)
	cache.sdMapLock.Lock()
	defer cache.sdMapLock.Unlock()
	v, ok := cache.sdMap[makeKey(symbol)]
	var sd *symbolData
	if !ok {
		return nil
	}
	sd = v
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

func (cache *XCache) getLastTickSerires(s *pb.Symbol) (*pb.TickSeries, error) {
	ctx := &api.Context{}
	req := &pb.ReqGetTickSeries{Symbol: s}
	ts, err := cache.ds.GetTickSerires(ctx, req)
	if err == nil {
		ret := &pb.TickSeries{List: ts.List}
		return ret, err
	}
	return nil, errors.New("empty")
}

func (cache *XCache) rGetKlineSeries(s *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) (*pb.KlineSeries, error) {
	ctx := &api.Context{}
	req := &pb.ReqGetKlineSeries{Symbol: s, Period: period, Start: startTime, End: endTime, LenLimit: lenLimit}
	var resp pb.KlineSeries
	ts, err := cache.ds.RGetKlineSeries(ctx, req)
	if err == nil {
		resp.Symbol = s
		resp.Period = period
		resp.List = ts.List
		return &resp, err
	}
	return nil, errors.New("empty")
}

func (cache *XCache) summary() *pb.CacheSummary {
	ret := &pb.CacheSummary{}
	cache.sdMapLock.Lock()
	defer cache.sdMapLock.Unlock()
	ret.KsMapSize = int32(len(cache.sdMap))
	for _, v := range cache.sdMap {
		sum := &pb.SymbolCacheSummary{}
		sum.Symbol = new(pb.Symbol)
		*sum.Symbol = v.Symbol
		sum.TickLen = int32(len(v.dayTS.List))
		ret.TotalTickLen += sum.TickLen
		for j := range v.KlineList {
			ks := &v.KlineList[j]
			sum.KlineLen += int32(len(ks.List))
		}
		ret.TotalKlineLen += sum.KlineLen
		ret.SymbolList = append(ret.SymbolList, sum)
		sum.SubscriberCount = int32(len(v.subscribers))
		ret.TotalSubscriberCount += sum.SubscriberCount
	}
	sort.Slice(ret.SymbolList, func(i, j int) bool {
		if ret.SymbolList[i].Symbol.Exchange < ret.SymbolList[j].Symbol.Exchange {
			return true
		}
		if ret.SymbolList[i].Symbol.Exchange > ret.SymbolList[j].Symbol.Exchange {
			return false
		}
		if ret.SymbolList[i].Symbol.Code < ret.SymbolList[j].Symbol.Code {
			return true
		}
		if ret.SymbolList[i].Symbol.Code > ret.SymbolList[j].Symbol.Code {
			return false
		}
		return false
	})
	return ret
}

func (cache *XCache) subscribe(req *pb.ReqSubscribe, ch chan *pb.MarketDataSnapshot) {
	for i := range req.List {
		sd := cache.getSymbolData(req.List[i])
		if sd != nil {
			sd.subscribe(ch)
		} else {
			log.Println("getSymbolData error ", sd.Symbol)
		}
	}
}

func (cache *XCache) unsubscribe(req *pb.ReqUnSubscribe, ch chan *pb.MarketDataSnapshot) {
	cache.sdMapLock.Lock()
	for _, v := range cache.sdMap {
		v.unsubscribe(ch)
	}
	cache.sdMapLock.Unlock()
}

func (cache *XCache) setTradingInstrument(req *pb.ReqSetTradingInstrument) {
	for i := range req.List {
		ti := req.List[i]
		cache.backend.SetTradingInstrument(ti)
		sd := cache.getSymbolData(ti.Symbol)
		if sd == nil {
			panic("sd==nil")
		}
		if sd != nil {
			sd.Lock()
			sd.ti = *ti
			sd.Unlock()
		}
	}
}

func (cache *XCache) getTradingInstrument(req *pb.ReqGetTradingInstrument) (*pb.TradingInstrument, error) {
	ctx := &api.Context{}
	ti, err := cache.ds.GetTradingInstrument(ctx, req.Symbol)
	if err == nil {
		return ti, err
	}
	return nil, errors.New("empty")
}

func (cache *XCache) getTick(s *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	ctx := &api.Context{}
	ti, err := cache.ds.GetLastTick(ctx, s)
	if err == nil {
		return ti, err
	}
	return nil, errors.New("empty")
}

func (cache *XCache) tradingInstrumentList(req *pb.ReqGetTradingInstrumentList) []*pb.TradingInstrument {
	ctx := &api.Context{}
	l, err := cache.ds.TradingInstrumentList(ctx, req)
	if err == nil {
		return l
	}
	return nil
}

package hub

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"

	pb "github.com/mineralres/goshare/pkg/pb/goshare"
	hubpb "github.com/mineralres/goshare/pkg/pb/hub"
)

type msg struct {
	t int
	d interface{}
}

// DemoEnv DemoEnv
type DemoEnv struct {
	demoOrderList     []*hubpb.DemoOrder
	demoOrderListLock sync.RWMutex
	chMsg             chan *msg
	mapTick           map[string]*pb.MarketDataSnapshot
	muMapTick         sync.RWMutex
	orderDB           *leveldb.DB
	options           *DemoEnvOptions
}

// DemoEnvOptions oiptions
type DemoEnvOptions struct {
	GetUID        func() string
	OnDemoOrder   func(*hubpb.DemoOrder)
	OnDemoTrade   func(*pb.Trade)
	GetInstrument func(string) *pb.Instrument
}

func isDone(do *hubpb.DemoOrder) bool {
	return do.VolumeTraded+do.VolumeCanceled == do.Volume || do.Status == int32(pb.OrderStatus_DONE)
}

// NewDemoEnv create demo env
func NewDemoEnv(options *DemoEnvOptions) *DemoEnv {
	ret := &DemoEnv{options: options}
	ret.mapTick = make(map[string]*pb.MarketDataSnapshot)
	ret.chMsg = make(chan *msg, 1000)
	var err error
	// 读取全部委托
	ret.orderDB, err = leveldb.OpenFile("demoorders", nil)
	if err != nil {
		panic("open level db error")
	}
	// 读取全部委托
	iter := ret.orderDB.NewIterator(util.BytesPrefix([]byte{}), nil)
	for iter.Next() {
		var do hubpb.DemoOrder
		err := proto.Unmarshal(iter.Value(), &do)
		if err != nil {
			continue
		}
		if isDone(&do) {
			ret.orderDB.Delete(iter.Key(), nil)
			continue
		}
		ret.demoOrderList = append(ret.demoOrderList, &do)
	}
	iter.Release()

	log.Println("[demo] 从缓存中读取", len(ret.demoOrderList), err)

	go func() {
		for m := range ret.chMsg {
			switch m.t {
			case 0:
				do := m.d.(hubpb.DemoOrder)
				key := fmt.Sprintf("%d-%d-%s", do.FrontId, do.SessionId, do.OrderRef)
				if isDone(&do) {
					log.Println("删除缓存", pb.OrderStatus(do.Status))
					ret.orderDB.Delete([]byte(key), nil)
				} else {
					d, err := proto.Marshal(&do)
					if err != nil {
						panic(err)
					}
					ret.orderDB.Put([]byte(key), d, nil)
					log.Println("保存到缓存", do.Symbol, pb.OrderStatus(do.Status))
				}
				options.OnDemoOrder(&do)
			case 1:
				tr := m.d.(pb.Trade)
				log.Println(tr)
				options.OnDemoTrade(&tr)
			}
		}
		for {
			select {}
		}
	}()
	return ret
}

func getMarketStatus(rule []*pb.MarketStatus) *pb.MarketStatus {
	t := time.Now().Unix() % 86400
	ret := &pb.MarketStatus{}
	for _, r := range rule {
		if t >= r.Time {
			ret = r
		}
	}
	return ret
}

// InsertDemoOrder 发送模拟
func (e *DemoEnv) InsertDemoOrder(req *hubpb.ReqInsertOrder) error {
	e.demoOrderListLock.Lock()
	defer e.demoOrderListLock.Unlock()
	for _, do := range e.demoOrderList {
		if do.FrontId == req.FrontId && do.SessionId == req.SessionId && do.OrderRef == req.OrderRef {
			return errors.New("禁止重复报单")
		}
	}
	inst := e.options.GetInstrument(req.Symbol)
	if inst == nil {
		return errors.New("[模拟]没有找到交易合约")
	}
	ms := getMarketStatus(inst.TimeRule)
	log.Println(req.Symbol, ms)
	if !ms.Send {
		return errors.New("当前状态禁止报单")
	}
	wd := time.Now().Weekday()
	if wd == time.Sunday {
		return errors.New("当前状态禁止报单")
	}
	if wd == time.Saturday {
		// 周六凌晨
		t := time.Now().Unix()
		if t%86400 > 3*3600 {
			return errors.New("当前状态禁止报单")
		}
	}
	log.Printf("[%s] 模拟发送委托,合约代码:[%s],类型[%d],数量[%d],价格[%.2f],交易日[%d]", req.TaId, req.Symbol, req.Offset, req.Volume, req.Price, inst.TradingDay)
	order := new(hubpb.DemoOrder)
	order.TaId = req.TaId
	order.BuId = req.BuId
	order.FrontId = req.FrontId
	order.SessionId = req.SessionId
	order.OrderRef = req.OrderRef
	order.Exchange = req.Exchange
	order.Symbol = req.Symbol
	order.Product = req.Product
	order.Direction = req.Direction
	order.Offset = req.Offset
	order.Price = req.Price
	order.Volume = req.Volume
	order.Product = inst.Product
	order.PriceTick = inst.PriceTick
	order.Multiple = inst.Multiple
	order.SendTradingDay = inst.TradingDay
	order.ProductType = inst.ProductType
	order.PriceType = req.PriceType
	order.TimeRule = inst.TimeRule
	order.SendTime = time.Now().Unix()
	order.DemoOrderId = e.options.GetUID()
	order.Status = int32(pb.OrderStatus_PENDING)
	order.MinLimitOrderVolume = inst.MinLimitOrderVolume
	order.MinMarketOrderVolume = inst.MinMarketOrderVolume

	e.chMsg <- &msg{d: *order}
	e.demoOrderList = append(e.demoOrderList, order)
	log.Println("InsertDemoOrder len ", len(e.demoOrderList))
	mds := e.getTick(req.Symbol)
	if mds != nil {
		go e.CheckDemoTrade(mds)
	}
	return nil
}

func (e *DemoEnv) getTick(symbol string) *pb.MarketDataSnapshot {
	e.muMapTick.RLock()
	defer e.muMapTick.RUnlock()
	v, ok := e.mapTick[symbol]
	if !ok {
		return nil
	}
	return v
}

// PushTick 收行情算成交
func (e *DemoEnv) PushTick(rtn *pb.MarketDataSnapshot) {
	e.muMapTick.Lock()
	e.mapTick[rtn.Symbol] = rtn
	e.muMapTick.Unlock()
	e.CheckDemoTrade(rtn)
}

// CheckDemoTrade 检查成交
func (e *DemoEnv) CheckDemoTrade(mds *pb.MarketDataSnapshot) {
	if len(mds.Depths) == 0 {
		log.Println("CheckDemoTrade", mds.Symbol, mds.Price, mds.Depths)
	}
	e.demoOrderListLock.RLock()
	defer e.demoOrderListLock.RUnlock()
	for _, do := range e.demoOrderList {
		if mds.Exchange == do.Exchange && mds.Symbol == do.Symbol {
			e.checkDemoOrderDone(mds, do)
		}
	}
}

func (e *DemoEnv) checkDemoOrderDone(mds *pb.MarketDataSnapshot, do *hubpb.DemoOrder) {
	order := do
	if order.SendTradingDay != mds.TradingDay {
		log.Println("tag1", order.SendTradingDay, mds.TradingDay)
		return
	}
	isCombine := order.ProductType == int32(pb.ProductType_COMBINATION)
	if order.OrderSourceType == int32(pb.OrderSourceType_COMBINATION_DERIVED) {
		// 套利衍生单不成交，而是直接用套利单计算成交
		return
	}
	if isDone(do) {
		return
	}
	nowT := time.Now().Unix()

	status := getMarketStatus(do.TimeRule)
	if !status.Matching {
		return
	}
	if mds.Depths[0].Bid > 0 && mds.Depths[0].Bid >= mds.UpperLimit && order.Direction == int32(pb.DirectionType_LONG) && mds.Depths[0].Bid < 1000000 {
		return
	}
	if mds.Depths[0].Ask > 0 && mds.Depths[0].Ask <= mds.LowerLimit && order.Direction == int32(pb.DirectionType_SHORT) && mds.Depths[0].Ask < 1000000 {
		return
	}
	if isCombine {
		if len(mds.Depths) == 0 {
			return
		}
		ob := mds.Depths[0]
		if ob.Ask <= 0 || ob.Ask >= 9999999 || ob.Bid <= 0 || ob.Bid >= 99999999 {
			return
		}
	} else {
		if order.Direction == int32(pb.DirectionType_LONG) && mds.Price > mds.UpperLimit-order.PriceTick {
			return
		}
		if order.Direction == int32(pb.DirectionType_SHORT) && mds.Price < mds.LowerLimit+order.PriceTick {
			return
		}
		if mds.Volume <= 0 {
			return
		}
		if mds.Price <= 0 || mds.Price > 99999999 {
			return
		}
	}
	bTraded := false
	ask := 0.0
	bid := 0.0
	for i := range mds.Depths {
		ob := mds.Depths[i]
		ask = ob.Ask
		if ob.AskVolume > 0 {
			break
		}
	}
	for i := range mds.Depths {
		ob := mds.Depths[i]
		bid = ob.Bid
		if ob.BidVolume > 0 {
			break
		}
	}
	lastPrice := mds.Price
	minimumVolume := do.MinLimitOrderVolume
	if order.Exchange == "SSE" || order.Exchange == "SZE" {
		minimumVolume = 100
	}
	isMarketOrder := false
	if order.PriceType == int32(pb.OrderPriceType_MARKET_PRICE) {
		minimumVolume = do.MinMarketOrderVolume
	}
	tradedAveragePrice := 0.0
	if isMarketOrder {
		bTraded = true
		if order.Direction == int32(pb.DirectionType_LONG) {
			tradedAveragePrice = ask
		} else {
			tradedAveragePrice = bid
		}
		if tradedAveragePrice > do.UpperLimit || tradedAveragePrice < do.LowerLimit {
			tradedAveragePrice = lastPrice
		}
	} else {
		if order.Direction == int32(pb.DirectionType_LONG) {
			if order.Price >= ask-0.000001 && ask > 0 && ask < 900000000 {
				tradedAveragePrice = ask
				bTraded = true
			}
		} else if order.Direction == int32(pb.DirectionType_SHORT) {
			if order.Price <= bid+0.0000001 && bid > 0 && bid < 900000000 {
				tradedAveragePrice = bid
				bTraded = true
			}
		}
	}
	ex := mds.Exchange
	tradedTime := mds.Time
	if (ex == "SSE" || ex == "SZE") && bTraded {
		// 股票故意延迟...
		if mds.Time-order.SendTime >= 1 {
			bTraded = true
		} else {
			bTraded = false
		}
		p := time.Unix(mds.Time, 0)
		h := p.Hour()
		m := p.Minute()
		if h == 9 && m >= 23 && m <= 26 {
			// 集合竞价成交价
			tradedAveragePrice = mds.Price
			tradedTime = time.Unix((mds.Time%86400)*86400, 0).Add(time.Hour * 9).Add(time.Minute * 25).Add(time.Second * 1).Unix()
		} else if mds.Exchange == "SZE" || ex == "SSE" {
			if (h == 14 && m >= 59) || (h == 15 && m <= 3) {
				tradedAveragePrice = mds.Price
				tradedTime = time.Unix((mds.Time%86400)*86400, 0).Add(time.Hour * 14).Add(time.Minute * 59).Add(time.Second * 59).Unix()
			}
		}
	}
	if tradedAveragePrice > 99990000 {
		bTraded = false
	}
	if bTraded {
		if tradedAveragePrice > mds.High && mds.High > 0 && mds.High < 99999999 {
			tradedAveragePrice = mds.High
		}
		if tradedAveragePrice < mds.Low && mds.Low > 0 && mds.Low < 99999999 {
			tradedAveragePrice = mds.Low
		}
		order.Status = int32(pb.OrderStatus_DONE)
		order.VolumeTraded = order.Volume
		e.chMsg <- &msg{d: *do}

		var tr pb.Trade
		tr.Direction = order.Direction
		tr.ExchangeOrderId = order.DemoOrderId
		tr.Price = tradedAveragePrice
		tr.Symbol = order.Symbol
		tr.TradedTime = tradedTime
		tr.TradedTradingDay = mds.TradingDay
		tr.TaId = order.TaId
		tr.BuId = order.BuId
		tr.Exchange = order.Exchange
		tr.FrontId = order.FrontId
		tr.SessionId = order.SessionId
		tr.OrderRef = order.OrderRef

		if ex == "CZCE" {
			tr.TradedTradingDay = do.TradingDay
		}
		tr.TradeType = int32(pb.TradeType_TT_NORMAL)
		if minimumVolume == 0 {
			log.Println("minimumVolue", minimumVolume)
			return
		}
		log.Printf("资管系统撮合成交[%.2f],合约:[%s], 方向:[%d], 开平:[%d],价格:[%.2f],数量:[%d],最小数量:[%d],time[%s] ask[%.4f] bid[%.4f]", tradedAveragePrice,
			tr.Symbol, tr.Direction, tr.Offset, tr.Price, tr.Volume, minimumVolume, time.Unix(mds.Time, 0).Format("15:04:05"), ask, bid)
		// 分多次成交
		v := order.Volume / minimumVolume
		mod := order.Volume % minimumVolume
		var tradedVolumes []int32
		lowerRange := 1
		if ex == "SSE" || ex == "SZE" {
			if order.Volume > 1000 {
				lowerRange = 2
			} else if order.Volume > 5000 {
				lowerRange = 3
			}
		} else if ex == "SHFE" || ex == "CZCE" || ex == "DCE" || ex == "CFFEX" || ex == "INE" {
			if order.Volume > 10 {
				lowerRange = 2
			} else if order.Volume > 30 {
				lowerRange = 3
			}
		}
		tradedVolumes = make([]int32, int(nowT%4)+lowerRange)
		for i := range tradedVolumes {
			tradedVolumes[i] = 0
		}
		rvLeft := v
		for {
			for i := range tradedVolumes {
				rv := int32(nowT % 3)
				if rvLeft < rv {
					rv = rvLeft
				}
				nowT++
				tradedVolumes[i] += rv
				rvLeft -= rv
				if rvLeft <= 0 {
					break
				}
			}
			if rvLeft <= 0 {
				break
			}
		}
		if mod > 0 {
			tradedVolumes = append(tradedVolumes, mod)
		}
		log.Println("generated volumes ", tradedVolumes)

		for i := range tradedVolumes {
			if tradedVolumes[i] > 0 {
				tr.Volume = tradedVolumes[i] * minimumVolume
				if tr.Volume > order.Volume {
					log.Println("Error volume order.volume = ", order.Volume, "trade.volume=", tr.Volume, "i=", i)
					tr.Volume = order.Volume
				}
				if isCombine {
					/*
						var s1, s2 pb.Symbol
						if base.ParseCombinationSymbol(&mds.Symbol, &s1, &s2) {
							var mds1, mds2 *pb.MarketDataSnapshot
							mds1 = e.te.GetLastTick(&s1)
							mds2 = e.te.GetLastTick(&s2)
							if mds1 != nil && mds2 != nil && err == nil {
								tr1 := tr
								tr1.Symbol = s1
								tr1.Price = mds1.Price
								tr2 := tr
								tr2.Symbol = s2
								tr2.Price = tr1.Price - tradedAveragePrice
								if tr1.Direction == int32(pb.DirectionType_LONG) {
									tr2.Direction = int32(pb.DirectionType_SHORT)
								} else {
									tr2.Direction = int32(pb.DirectionType_LONG)
								}
								// 涨跌停检查和调整.
								tr1.TradeId, _ = e.redisClient.GetUniqueString()
								tr2.TradeId, _ = e.redisClient.GetUniqueString()
								e.trChan <- tr1
								e.trChan <- tr2
							}
						}
					*/
				} else {
					tr.TradeId = e.options.GetUID()
					rtn := tr
					e.chMsg <- &msg{t: 1, d: rtn}
					log.Printf("DemoUTC trade index[%d], volume[%d], isCombine[%t], tradedTime[%d], mdsTime[%d] symbol[%s] tradeid[%s]", i, tr.Volume,
						isCombine, tradedTime, mds.Time, mds.Symbol, tr.TradeId)
				}
			}
		}
	}
}

// CancelDemoOrder cancel demo order
func (e *DemoEnv) CancelDemoOrder(req *hubpb.ReqCancelOrder) error {
	e.demoOrderListLock.Lock()
	defer e.demoOrderListLock.Unlock()
	for _, do := range e.demoOrderList {
		order := do
		if req.FrontId == do.FrontId && req.SessionId == do.SessionId && req.OrderRef == do.OrderRef {
			if isDone(do) {
				return errors.New("当前状态不能撤单")
			}
			ms := getMarketStatus(do.TimeRule)
			log.Println(req.Symbol, ms)
			if !ms.Cancel {
				log.Println("CancelDemoOrder 当前状态不能撤单", do.DemoOrderId)
				return errors.New("当前状态不能撤单")
			}
			order.Status = int32(pb.OrderStatus_CANCELED)
			order.VolumeCanceled = order.Volume - order.VolumeTraded
			log.Println("CancelDemoOrder 撤单成功", do.DemoOrderId)
			e.chMsg <- &msg{d: *do}
			return nil
		}
	}
	return errors.New("notfound")
}

// CurrentDemoOrderList do list
func (e *DemoEnv) CurrentDemoOrderList() []hubpb.DemoOrder {
	var ret []hubpb.DemoOrder
	e.demoOrderListLock.Lock()
	defer e.demoOrderListLock.Unlock()
	for _, do := range e.demoOrderList {
		if !isDone(do) {
			ret = append(ret, *do)
		}
	}
	return ret
}

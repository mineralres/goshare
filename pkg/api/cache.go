package api

import (
	"log"
	"strconv"
	"time"

	"github.com/mineralres/goshare/pkg/pb"
	"github.com/syndtr/goleveldb/leveldb"
)

// Cache 数据聚合
type Cache struct {
	dsList       []DataSource          // 常规数据
	realtimeList []RealtimeDataSource  // 实时订阅
	um           UserManager           // 用户管理
	sm           StrategyManager       // 策略管理
	tam          TradingAccountManager // 交易账户管理
	backend      CacheBackend          // 本地缓存存储
}

// CacheBackend 缓存
type CacheBackend interface {
	Set(key, val []byte) error
	Get(key []byte) ([]byte, error)
}

type ldbBackend struct {
	xdb *leveldb.DB
}

func (lb *ldbBackend) Set(key, val []byte) error {
	return lb.xdb.Put(key, val, nil)
}

func (lb *ldbBackend) Get(key []byte) ([]byte, error) {
	return lb.xdb.Get(key, nil)
}

// NewCache NewCache2
func NewCache(dsList []DataSource, realtimeList []RealtimeDataSource, um UserManager, sm StrategyManager, tam TradingAccountManager) *Cache {
	ret := &Cache{dsList: dsList, realtimeList: realtimeList, um: um, sm: sm, tam: tam}
	bn := &ldbBackend{}
	var err error
	bn.xdb, err = leveldb.OpenFile("db/apicache", nil)
	if err != nil || bn.xdb == nil {
		log.Println(err)
		panic("open level db error")
	}
	ret.backend = bn
	return ret
}

func (c *Cache) subscribe(req *pb.ReqSubscribe, ch chan *pb.MarketDataSnapshot) {
	for i := range c.realtimeList {
		if _, err := c.realtimeList[i].Subscribe(&Context{}, req, ch); err == nil {
			return
		}
	}
}

func (c *Cache) unsubscribe(req *pb.ReqUnSubscribe, ch chan *pb.MarketDataSnapshot) {
}

func (c *Cache) getMainContract() ([]*pb.TradingInstrument, error) {
	var ctx Context
	var resp []*pb.TradingInstrument
	var err error
	for _, ds := range c.dsList {
		if resp, err = ds.TradingInstrumentList(&ctx, &pb.ReqGetTradingInstrumentList{}); err == nil {
			break
		}
	}
	m := make(map[string]*pb.TradingInstrument)
	for _, ti := range resp {
		v, ok := m[ti.ProductInfo.ProductId.Code]
		if ok {
			if ti.InstrumentInfo.PrePosition >= v.InstrumentInfo.PrePosition {
				m[ti.ProductInfo.ProductId.Code] = ti
			}
		} else {
			m[ti.ProductInfo.ProductId.Code] = ti
		}
	}
	var ret []*pb.TradingInstrument
	for _, ti := range m {
		ret = append(ret, ti)
	}
	return ret, nil
}

func getDay() int32 {
	str := time.Now().Format("20060102")
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return int32(i)
}

package api

import "github.com/mineralres/goshare/pkg/pb"

// Cache 数据聚合
type Cache struct {
	dsList       []DataSource          // 常规数据
	realtimeList []RealtimeDataSource  // 实时订阅
	um           UserManager           // 用户管理
	sm           StrategyManager       // 策略管理
	tam          TradingAccountManager // 交易账户管理
}

// NewCache create new cache
func NewCache() *Cache {
	ret := new(Cache)
	return ret
}

// NewCache2 NewCache2
func NewCache2(dsList []DataSource, realtimeList []RealtimeDataSource, um UserManager, sm StrategyManager, tam TradingAccountManager) *Cache {
	ret := &Cache{dsList: dsList, realtimeList: realtimeList, um: um, sm: sm, tam: tam}
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
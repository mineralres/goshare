package api

import (
	pb "github.com/mineralres/protos/src/go/goshare"
)

// GoshareAPI 平台核心
type GoshareAPI interface {
	DataSource
	RealtimeDataSource
	UserManager
	StrategyManager
	TradingAccountManager
}

// Context Context
type Context struct {
}

// DataSource 数据源
type DataSource interface {
	// 读取tick
	GetLastTick(*Context, string, string) (*pb.MarketDataSnapshot, error)
	// 取最新一天的tick序列
	GetTickSerires(*Context, *pb.ReqGetTickSeries) (*pb.RspGetTickSeries, error)
	// 读取合约信息
	GetTradingInstrument(*Context, string, string) (*pb.Instrument, error)
	// 反向取K线
	RGetKlineSeries(*Context, *pb.ReqGetKlineSeries) (*pb.RspGetKlineSeries, error)
	// K线
	GetKlineSeries(*Context, *pb.ReqGetKlineSeries) (*pb.RspGetKlineSeries, error)
}

// RealtimeDataSource 实时推送数据源
type RealtimeDataSource interface {
	Subscribe(*Context, *pb.ReqSubscribe, chan *pb.MarketDataSnapshot) (*pb.RspSubscribe, error)
	UnSubscribe(*Context, *pb.ReqUnSubscribe, chan *pb.MarketDataSnapshot) (*pb.RspUnSubscribe, error)
}

// UserManager 用户管理
type UserManager interface {
}

// StrategyManager 策略管理
type StrategyManager interface {
}

// TradingAccountManager 交易通道管理
type TradingAccountManager interface {
}

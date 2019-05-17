package ta

import "github.com/mineralres/goshare/pkg/pb"

// 交易通道

// Account 交易通道
type Account interface {
	Init() error                              //初始化
	Login(*pb.TradingAccount) error           // 登陆
	Disconnect() error                        // 断开
	InsertOrder(*pb.Order) error              // 发单
	CancelOrder(*pb.CancelOrderRequest) error // 撤单
}

// Handler 回调
type Handler interface {
	OnFrontConnected()                         // 连接上
	OnRspUserLogin(*pb.RspTradingAccountLogin) // 登陆返回
	OnRtnOrder(*pb.Order)                      // 委托返回
	OnRtnTrade(*pb.TradeReport)                // 成交返回
	OnRspOrderInsert(*pb.RspOrderInsert)       // 发单失败
	OnOther(pb.BindingMessageType, []byte)     //  其他
}

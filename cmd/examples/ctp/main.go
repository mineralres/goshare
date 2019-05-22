package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/mineralres/goshare/pkg/util"
	"github.com/mineralres/goshare/pkg/pb"
	ta "github.com/mineralres/goshare/pkg/trading_account"
)

type mdhandler struct {
	mdAPI *ta.MarketDataAPI
}

func (h *mdhandler) OnFrontConnected() {
	log.Println("行情前置连接成功")
	h.mdAPI.Login(&pb.TradingAccount{})
}

func (h *mdhandler) OnRspUserLogin(rsp *pb.RspTradingAccountLogin) {
	log.Println("行情前置返回:", util.StringFromGBK2(rsp.ErrorMsg))
	var l pb.SymbolList
	l.List = append(l.List, &symbol)
	h.mdAPI.Subscribe(&l)
}

func (h *mdhandler) OnRtnDepthMarketData(md *pb.MarketDataSnapshot) {
	m := jsonpb.Marshaler{}
	d, err := m.MarshalToString(md)
	log.Println("行情返回", string(d), err)
}

// 交易回调
type myhandler struct {
	api ta.Account
}

func (h *myhandler) OnRtnOrder(rtn *pb.Order) {
	log.Printf("[OnRtnOrder]委托回报 ID[%d %d %d] [%s] LimitPrice[%.4f] Comment[%s]", rtn.Id.FrontId, rtn.Id.SessionId, rtn.Id.OrderRef, rtn.Symbol.Code, rtn.LimitPrice, rtn.Comment)
}

func (h *myhandler) OnRtnTrade(rtn *pb.TradeReport) {
	log.Printf("[OnRtnTrade]成交回报 [%d][%s] Price[%.4f]", rtn.Symbol.Exchange, rtn.Symbol.Code, rtn.Price)
}

func (h *myhandler) OnRtnCancelOrder(*pb.OnRtnCancelOrder) {

}

func (h *myhandler) OnRspUserLogin(rsp *pb.RspTradingAccountLogin) {
	log.Println("交易前置登陆返回:", *rsp, util.StringFromGBK2(rsp.ErrorMsg))
	order := pb.Order{}
	order.Id = &pb.OrderID{OrderRef: 1}
	order.Symbol = &symbol
	order.Direction = pb.OrderDirection_OD_LONG
	order.OffsetFlag = pb.OffsetFlag_OF_OPEN
	order.LimitPrice = 11510
	order.Volume = 1
	order.Account = account
	ret := h.api.InsertOrder(&order)
	log.Println("发单返回 ", ret)
	time.Sleep(time.Second)
	reqCancel := &pb.CancelOrderRequest{}
	reqCancel.OrderId = &pb.OrderID{}
	reqCancel.Account = order.Account
	reqCancel.OrderId.FrontId = rsp.FrontId
	reqCancel.OrderId.SessionId = rsp.SessionId
	reqCancel.OrderId.OrderRef = order.Id.OrderRef
	reqCancel.Symbol = order.Symbol
	h.api.CancelOrder(reqCancel)
}

func (h *myhandler) OnRspOrderInsert(rsp *pb.RspOrderInsert) {
	log.Println("发单失败", util.StringFromGBK2(rsp.ErrorMsg), rsp.OrderRef)
}

func (h *myhandler) OnFrontConnected() {
	a := pb.TradingAccount{Account: account, Password: password}
	log.Println("交易前置连接成功")
	h.api.Login(&a)
}

var (
	account  = "xxx" // simnow 账号
	password = "xxx" // simnow 密码
	symbol   = pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "ru1909"}
	r        = &pb.TradingRoute{BrokerId: "9999", MarketDataFrontList: []string{"180.169.99.50:41213"}, TradingFrontList: []string{"180.168.146.187:10000"}}
)

func testTrading() {
	log.Println("测试交易")
	var h myhandler
	h.api = ta.MakeTrader(r, &h)
	h.api.Init()
}

func testMarketData() {
	log.Println("测试行情")
	var mh mdhandler
	mh.mdAPI = ta.MakeMarketDataAPI(r, &mh)
	mh.mdAPI.Init()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	testMarketData()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
}

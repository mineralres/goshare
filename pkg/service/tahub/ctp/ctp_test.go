package ctp

import (
	"log"
	"testing"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/mineralres/goshare/pkg/util"
	"github.com/mineralres/goshare/pkg/pb"
)

type mdhandler struct {
	mdAPI *MarketDataAPI
}

func (h *mdhandler) OnFrontConnected() {
	var ta pb.TradingAccount
	log.Println("MD 连接成功")
	h.mdAPI.Login(&ta)
}

func (h *mdhandler) OnRspUserLogin(rsp *pb.RspTradingAccountLogin) {
	log.Println("OnRspUserLogin", *rsp, util.StringFromGBK2(rsp.ErrorMsg))
	var l pb.SymbolList
	l.List = append(l.List, &pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "ru1909"})
	h.mdAPI.Subscribe(&l)
}

func (h *mdhandler) OnRtnDepthMarketData(md *pb.MarketDataSnapshot) {
	m := jsonpb.Marshaler{}
	d, err := m.MarshalToString(md)
	log.Println(string(d), err)
}

type myhandler struct {
	api Account
}

func (h *myhandler) OnRtnOrder(rtn *pb.Order) {
	log.Printf("OnRtnOrder ID[%d %d %d] [%s] LimitPrice[%.4f] Comment[%s]", rtn.Id.FrontId, rtn.Id.SessionId, rtn.Id.OrderRef, rtn.Symbol.Code, rtn.LimitPrice, rtn.Comment)
}

func (h *myhandler) OnRtnTrade(rtn *pb.TradeReport) {
	log.Printf("OnRtnTrade [%d][%s] Price[%.4f]", rtn.Symbol.Exchange, rtn.Symbol.Code, rtn.Price)
}

func (h *myhandler) OnRtnCancelOrder(*pb.OnRtnCancelOrder) {

}

func (h *myhandler) OnRspUserLogin(rsp *pb.RspTradingAccountLogin) {
	log.Println("OnRspUserLogin", *rsp, util.StringFromGBK2(rsp.ErrorMsg))
	order := pb.Order{}
	order.Id = &pb.OrderID{OrderRef: 1}
	order.Symbol = &pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "ru1909"}
	order.Direction = pb.OrderDirection_OD_LONG
	order.OffsetFlag = pb.OffsetFlag_OF_OPEN
	order.LimitPrice = 11510
	order.Volume = 1
	order.Account = "059926"
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
	var ta pb.TradingAccount
	ta.Account = ""  // simnow account
	ta.Password = "" // simnow password
	log.Println("OnFrontConnected")
	h.api.Login(&ta)
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestDll(t *testing.T) {
	r := &pb.TradingRoute{}
	r.BrokerId = "9999"
	r.MarketDataFrontList = append(r.MarketDataFrontList, "180.169.99.50:41213")
	r.TradingFrontList = append(r.TradingFrontList, "180.168.146.187:10000")

	// var mh mdhandler
	// mh.mdAPI = MakeMarketDataAPI(r, &mh)
	// mh.mdAPI.Init()

	var h myhandler
	h.api = MakeTrader(r, &h)
	h.api.Init()

	time.Sleep(time.Second * 30)

}

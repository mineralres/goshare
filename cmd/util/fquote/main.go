package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/golang/protobuf/proto"
	"github.com/mineralres/goshare/pkg/base"
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/service"
	"github.com/mineralres/goshare/pkg/ta/ctp"
)

type mdhandler struct {
	mdAPI      *ctp.MarketDataAPI
	symbolList []*pb.Symbol
	cl         pb.DCenterClient
	tiMap      sync.Map
	tiList     []*pb.TradingInstrument
}

func (h *mdhandler) OnFrontConnected() {
	log.Println("行情前置连接成功")
	h.mdAPI.Login(&pb.TradingAccount{})
}

func (h *mdhandler) OnRspUserLogin(rsp *pb.RspTradingAccountLogin) {
	log.Println("行情前置返回:", base.StringFromGBK2(rsp.ErrorMsg))
	conn, err := service.GetClientConn("srv.dcenter")
	if err != nil {
		panic(err)
	}
	h.cl = pb.NewDCenterClient(conn)

	h.cl.SetTradingInstrument(context.Background(), &pb.ReqSetTradingInstrument{List: h.tiList})

	var l pb.SymbolList
	l.List = h.symbolList
	for _, ti := range h.tiList {
		h.tiMap.Store(ti.Symbol.Code, ti)
	}
	h.mdAPI.Subscribe(&l)
}

func (h *mdhandler) OnRtnDepthMarketData(md *pb.MarketDataSnapshot) {
	v, ok := h.tiMap.Load(md.Symbol.Code)
	if ok {
		ti := v.(*pb.TradingInstrument)
		md.Symbol.Exchange = ti.Symbol.Exchange
		md.PriceTick = ti.ProductInfo.PriceTick
		md.VolumeMultiple = ti.ProductInfo.VolumeMultiple
		md.Name = ti.InstrumentInfo.SymbolName
	}
	resp, err := h.cl.UpdateTick(context.Background(), md)
	if err != nil {
		// h.cl = pb.NewDataSourceService("go.micro.srv.mddb", grpc.NewService().Client())
		log.Println(md.Symbol, resp, err)
	}
}

// 交易回调
type myhandler struct {
	api *ctp.Trader
	mdh mdhandler
}

func (h *myhandler) OnRtnOrder(rtn *pb.Order) {
	log.Printf("[OnRtnOrder]委托回报 ID[%d %d %d] [%s] LimitPrice[%.4f] Comment[%s]", rtn.Id.FrontId, rtn.Id.SessionId, rtn.Id.OrderRef, rtn.Symbol.Code, rtn.LimitPrice, rtn.Comment)
}

func (h *myhandler) OnRtnTrade(rtn *pb.TradeReport) {
	log.Printf("[OnRtnTrade]成交回报 [%d][%s] Price[%.4f]", rtn.Symbol.Exchange, rtn.Symbol.Code, rtn.Price)
}

func (h *myhandler) OnRtnCancelOrder(*pb.OnRtnCancelOrder) {

}

func (h *myhandler) OnOther(t pb.BindingMessageType, d []byte) {
	switch t {
	case pb.BindingMessageType_CTP_ON_RSP_QRY_INSTRUMENT:
		var rsp pb.CTPOnRspQryInstrument
		if err := proto.Unmarshal(d, &rsp); err == nil {
			ti := ctp.FromCTPInstrumentField(rsp.Inst)
			// log.Println("查询合约返回", ti.Symbol.Code, strings.Contains(ti.Symbol.Code, "efp"))
			if !strings.Contains(ti.Symbol.Code, "efp") {
				h.mdh.symbolList = append(h.mdh.symbolList, ti.Symbol)
				h.mdh.tiList = append(h.mdh.tiList, ti)
			}
			if rsp.Response.IsLast {
				h.mdh.mdAPI = ctp.MakeMarketDataAPI(r, &h.mdh)
				h.mdh.mdAPI.Init()
			}

		}
	}
}

func (h *myhandler) OnRspUserLogin(rsp *pb.RspTradingAccountLogin) {
	log.Println("交易前置登陆返回:", *rsp, base.StringFromGBK2(rsp.ErrorMsg))
	h.api.ReqQryInstrument(&pb.Symbol{})
}

func (h *myhandler) OnRspOrderInsert(rsp *pb.RspOrderInsert) {
	log.Println("发单失败", base.StringFromGBK2(rsp.ErrorMsg), rsp.OrderRef)
}

func (h *myhandler) OnFrontConnected() {
	a := pb.TradingAccount{Account: account, Password: password}
	log.Println("交易前置连接成功")
	h.api.Login(&a)
}

var (
	account  = "059926" // simnow 账号
	password = "xxxxxx" // simnow 密码
	symbol   = pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "ru1909"}
	r        = &pb.TradingRoute{BrokerId: "9999", MarketDataFrontList: []string{"180.169.99.50:41213"}, TradingFrontList: []string{"180.168.146.187:10000"}}
)

func testTrading() {
}

func testMarketData() {
	log.Println("测试行情")
	var mh mdhandler
	mh.mdAPI = ctp.MakeMarketDataAPI(r, &mh)
	mh.mdAPI.Init()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var h myhandler
	h.api = ctp.MakeTrader(r, &h)
	h.api.Init()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
}

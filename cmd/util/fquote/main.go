package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/golang/protobuf/proto"
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/ta/ctp"
	"github.com/mineralres/goshare/pkg/util"
	"github.com/mineralres/goshare/pkg/util/datasource"
)

type mdhandler struct {
	mdAPI      *ctp.MarketDataAPI
	symbolList []*pb.Symbol
	tiMap      sync.Map
	tiList     []*pb.TradingInstrument
	c          config
	cl         *datasource.Client
}

func (h *mdhandler) OnFrontConnected() {
	log.Println("行情前置连接成功")
	h.mdAPI.Login(&pb.TradingAccount{})
}

func (h *mdhandler) OnRspUserLogin(rsp *pb.RspTradingAccountLogin) {
	log.Println("行情前置返回:", util.StringFromGBK2(rsp.ErrorMsg))
	options := &datasource.ClientOptions{}
	options.URL.Scheme = h.c.Scheme
	options.URL.Host = h.c.Host
	options.Token = h.c.Token
	h.cl = datasource.MakeClient(options)
	h.cl.SetTradingInstrument(&pb.ReqSetTradingInstrument{List: h.tiList})

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
	err := h.cl.UpdateTick(md)
	if err != nil {
		log.Println(md.Symbol, err)
	}
}

// 交易回调
type tradehandler struct {
	api *ctp.Trader
	mdh mdhandler
	c   config
}

func (h *tradehandler) OnRtnOrder(rtn *pb.Order) {
	log.Printf("[OnRtnOrder]委托回报 ID[%d %d %d] [%s] LimitPrice[%.4f] Comment[%s]", rtn.Id.FrontId, rtn.Id.SessionId, rtn.Id.OrderRef, rtn.Symbol.Code, rtn.LimitPrice, rtn.Comment)
}

func (h *tradehandler) OnRtnTrade(rtn *pb.TradeReport) {
	log.Printf("[OnRtnTrade]成交回报 [%d][%s] Price[%.4f]", rtn.Symbol.Exchange, rtn.Symbol.Code, rtn.Price)
}

func (h *tradehandler) OnRtnCancelOrder(*pb.OnRtnCancelOrder) {

}

func (h *tradehandler) OnOther(t pb.BindingMessageType, d []byte) {
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
				h.mdh.c = h.c
				r := &pb.TradingRoute{BrokerId: h.c.BrokerID, MarketDataFrontList: h.c.MarketDataFrontList, TradingFrontList: h.c.TradeFrontList}
				h.mdh.mdAPI = ctp.MakeMarketDataAPI(r, &h.mdh)
				h.mdh.mdAPI.Init()
			}

		}
	}
}

func (h *tradehandler) OnRspUserLogin(rsp *pb.RspTradingAccountLogin) {
	log.Println("交易前置登陆返回:", *rsp, util.StringFromGBK2(rsp.ErrorMsg))
	h.api.ReqQryInstrument(&pb.Symbol{})
}

func (h *tradehandler) OnRspOrderInsert(rsp *pb.RspOrderInsert) {
	log.Println("发单失败", util.StringFromGBK2(rsp.ErrorMsg), rsp.OrderRef)
}

func (h *tradehandler) OnFrontConnected() {
	a := pb.TradingAccount{Account: h.c.Account, Password: h.c.Password}
	log.Println("交易前置连接成功")
	h.api.Login(&a)
}

var (
// account  = "059926" // simnow 账号
// password = "198612" // simnow 密码
// symbol   = pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "ru1909"}
// r        = &pb.TradingRoute{BrokerId: "9999", MarketDataFrontList: []string{"180.169.99.50:41213"}, TradingFrontList: []string{"180.168.146.187:10000"}}
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var c config
	err := loadConfig("config.json", &c)
	if err != nil {
		panic(err)
	}
	log.Println("config", c)

	r := &pb.TradingRoute{BrokerId: c.BrokerID, MarketDataFrontList: c.MarketDataFrontList, TradingFrontList: c.TradeFrontList}

	var h tradehandler
	h.c = c
	h.api = ctp.MakeTrader(r, &h)
	h.api.Init()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
}

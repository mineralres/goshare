package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/mineralres/goshare/pkg/pb/goshare"
	"github.com/mineralres/goshare/pkg/service/tahub/ctp"
	"github.com/mineralres/goshare/pkg/util"
	"github.com/mineralres/goshare/pkg/util/datasource"
)

type mdhandler struct {
	mdAPI         *ctp.MarketDataAPI
	symbolList    []*pb.Symbol
	tiMap         map[string]*pb.TradingInstrument
	tiList        []*pb.TradingInstrument
	c             config
	cl            *datasource.Client
	nPriceUpdated int64
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
	err := h.cl.SetTradingInstrument(&pb.ReqSetTradingInstrument{List: h.tiList})
	log.Printf("上传合约 [%d] 个, %v", len(h.tiList), err)

	var l pb.SymbolList
	l.List = h.symbolList
	h.tiMap = make(map[string]*pb.TradingInstrument)
	for _, ti := range h.tiList {
		h.tiMap[ti.Symbol.Code] = ti
	}
	h.mdAPI.Subscribe(&l)
}

func (h *mdhandler) OnRtnDepthMarketData(md *pb.MarketDataSnapshot) {
	ti, ok := h.tiMap[md.Symbol.Code]
	if ok {
		if ti.InstrumentInfo.PrePosition == 0 {
			h.nPriceUpdated++
			ti.InstrumentInfo.PrePosition = md.PrePosition
			ti.InstrumentInfo.UpperLimitPrice = md.UpperLimitPrice
			ti.InstrumentInfo.LowerLimitPrice = md.LowerLimitPrice
			ti.InstrumentInfo.PreClosePrice = md.PreClose
			ti.InstrumentInfo.PreSettlementPrice = md.PreSettlementPrice
			ti.InstrumentInfo.SettlementPrice = md.SettlementPrice
			ti.InstrumentInfo.UpdateTradingDay = md.TradingDay
			ti.InstrumentInfo.UpdateTime = md.Time
			if int(h.nPriceUpdated) == len(h.tiMap) {
				// 完成价格修补再上传合约
				var l []*pb.TradingInstrument
				for _, ti := range h.tiMap {
					l = append(l, ti)
				}
				err := h.cl.SetTradingInstrument(&pb.ReqSetTradingInstrument{List: l})
				log.Printf("二次上传合约 [%d] 个, %v", len(h.tiList), err)
			}
		}
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

	r := &pb.TradingRoute{BrokerId: c.BrokerID, MarketDataFrontList: c.MarketDataFrontList, TradingFrontList: c.TradeFrontList}

	var h tradehandler
	h.c = c
	h.api = ctp.MakeTrader(r, &h)
	h.api.Init()

	start := flag.String("start", "", "exit when timeout")
	end := flag.String("end", "", "exit when timeout")
	flag.Parse()

	ticker := time.NewTicker(time.Second)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	for {
		select {
		case sig := <-ch:
			log.Println("exit on signal ", sig)
			return
		case <-ticker.C:
			if len(*start) > 0 && len(*end) > 0 {
				now := time.Now().Format("15:04:05")
				// log.Println("now", now, *start, *end)
				if *start <= *end && now >= *start && now <= *end {
					// 白盘
				} else if *start > *end && !(now >= *end && now <= *start) {
					// 夜盘
				} else {
					log.Printf("exit on now[%s] start[%s] end[%s]", now, *start, *end)
					return
				}
			}
		}
	}

}

package tradingaccount

import (
	"log"
	"testing"
	"time"

	"github.com/mineralres/goshare/pkg/base"
	"github.com/mineralres/goshare/pkg/pb"
)

type mdhandler struct {
	mdAPI *MarketDataAPI
}

func (h *mdhandler) OnRspUserLogin(rsp *pb.CTPRspInfo) {
	log.Println("OnRspUserLogin", *rsp, base.StringFromGBK2(rsp.ErrorMsg))
	var l pb.SymbolList
	l.List = append(l.List, &pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "ru1905"})
	h.mdAPI.Subscribe(&l)
}

func (h *mdhandler) OnRtnDepthMarketData(md *pb.MarketDataSnapshot) {
	log.Printf("OnRtnDepthMarketData[%s] [%.4f]", md.Symbol.Code, md.Price)
}

type myhandler struct {
	api Account
}

func (h *myhandler) OnRtnOrder(*pb.Order) {

}

func (h *myhandler) OnRtnTrade(*pb.TradeReport) {

}

func (h *myhandler) OnRtnCancelOrder(*pb.OnRtnCancelOrder) {

}

func (h *myhandler) OnRspUserLogin(rsp *pb.CTPRspInfo) {
	log.Println("OnRspUserLogin", *rsp, base.StringFromGBK2(rsp.ErrorMsg))
	order := pb.Order{}
	order.Id = &pb.OrderID{OrderRef: 1}
	order.Symbol = &pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "ru1905"}
	order.Direction = pb.OrderDirection_OD_LONG
	order.OffsetFlag = pb.OffsetFlag_OF_OPEN
	order.LimitPrice = 11395
	order.Volume = 1
	h.api.InsertOrder(&order)
}

func (h *myhandler) OnRspOrderInsert(rsp *pb.CTPOnRspOrderInsert) {
	log.Println("发单失败", base.StringFromGBK2(rsp.ErrorMsg), rsp.OrderRef)
}

func (h *myhandler) OnFrontConnected() {
	var ta pb.TradingAccount
	// ta.Account = "059926"
	// ta.Password = "198612"
	ta.Account = "15900868219"
	ta.Password = "123456"
	log.Println("OnFrontConnected")
	h.api.Login(&ta)
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestDll(t *testing.T) {
	// cmd := exec.Command("explorer", "https://www.baidu.com")
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	r := &pb.TradingRoute{}
	r.BrokerId = "ta"
	// r.MarketDataFrontList = append(r.MarketDataFrontList, "180.169.99.50:41213")
	// r.MarketDataFrontList = append(r.MarketDataFrontList, "180.168.146.187:10010")
	r.MarketDataFrontList = append(r.MarketDataFrontList, "localhost:41205")

	var h myhandler
	// h.mdAPI = MakeMarketDataAPI(r, &h)
	// h.mdAPI.Init()

	h.api = MakeTrader(r, &h)
	h.api.Init()

	time.Sleep(time.Second * 30)

}

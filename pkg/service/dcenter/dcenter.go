package dcenter

import (
	"context"
	"log"

	"github.com/mineralres/goshare/pkg/pb"
)

// DBEngine 数据存储引擎
type DBEngine interface {
	// 保存单根K线
	Save(*pb.Symbol, pb.PeriodType, *pb.Kline) error
	// 保存日tick序列
	SaveDayTickSeries(ts *pb.TickSeries) error
	// 反向取K线
	RGetKlineSeries(s *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) *pb.KlineSeries
	// 取最新一天的tick序列
	GetLastTickSerires(s *pb.Symbol) *pb.TickSeries
	// 保存合约信息
	SetTradingInstrument(*pb.TradingInstrument) error
	// 读取合约信息
	GetTradingInstrument(*pb.Symbol) (*pb.TradingInstrument, error)
	// 全部合约信息
	TradingInstrumentList() []*pb.TradingInstrument
	// 保存tick
	SetTick(*pb.MarketDataSnapshot) error
	// 读取tick
	GetTick(*pb.Symbol) (*pb.MarketDataSnapshot, error)
}

// RPCHandler RPCHandler
type RPCHandler struct {
	cache *XCache
}

// MakeRPCHandler 服务
func MakeRPCHandler() *RPCHandler {
	h := &RPCHandler{}
	h.cache = MakeXCache()
	return h
}

// SaveKline SaveKline
func (h *RPCHandler) SaveKline(ctx context.Context, req *pb.ReqSaveKline) (*pb.RspSaveKline, error) {
	return nil, nil
}

// Subscribe 订阅行情
func (h *RPCHandler) Subscribe(req *pb.ReqSubscribe, stream pb.DCenter_SubscribeServer) error {
	log.Println("DCenter Subscribe", req)
	ch := make(chan pb.MarketDataSnapshot, 9999) // fixme: 这个数量需要调整
	h.cache.subscribe(req, ch)
	defer h.cache.unsubscribe(req, ch)
	for {
		select {
		case md := <-ch:
			err := stream.Send(&md)
			if err != nil {
				log.Println("Subscribe err ", err)
				break
			}
		case <-stream.Context().Done():
			log.Println("Context().Done()")
			return nil
		}
	}
	return nil
}

// CombineSubscribe 综合订阅
func (h *RPCHandler) CombineSubscribe(req *pb.ReqCombineSubscribe, stream pb.DCenter_CombineSubscribeServer) error {
	return nil
}

// UpdateTick 行情推送更新
func (h *RPCHandler) UpdateTick(ctx context.Context, req *pb.MarketDataSnapshot) (*pb.EmptyResponse, error) {
	if req.Symbol == nil {
		req.Symbol = &pb.Symbol{}
	}
	// log.Printf("UpdateTick [%s] price[%.4f] depth[%d] time[%d]", req.Symbol.Code, req.Price, len(req.OrderBookList), req.Time)
	h.cache.Update(req)
	return &pb.EmptyResponse{}, nil
}

// GetDCenterInfo GetDCenterInfo
func (h *RPCHandler) GetDCenterInfo(ctx context.Context, req *pb.ReqGetDCenterInfo) (*pb.RspGetDCenterInfo, error) {
	res := &pb.RspGetDCenterInfo{}
	res.CacheSummary = h.cache.summary()
	return res, nil
}

// GetTradingInstrument GetTradingInstrument
func (h *RPCHandler) GetTradingInstrument(ctx context.Context, req *pb.ReqGetTradingInstrument) (*pb.RspGetTradingInstrument, error) {
	var ret pb.RspGetTradingInstrument
	return &ret, nil
}

// SetTradingInstrument SetTradingInstrument
func (h *RPCHandler) SetTradingInstrument(ctx context.Context, req *pb.ReqSetTradingInstrument) (*pb.RspSetTradingInstrument, error) {
	var ret pb.RspSetTradingInstrument
	h.cache.setTradingInsrument(req)
	return &ret, nil
}

// GetTradingInstrumentList ReqGetTradingInstrumentList
func (h *RPCHandler) GetTradingInstrumentList(ctx context.Context, req *pb.ReqGetTradingInstrumentList) (*pb.RspGetTradingInstrumentList, error) {
	var ret pb.RspGetTradingInstrumentList
	return &ret, nil
}

package dcenter

import (
	"context"
	"log"

	"github.com/mineralres/goshare/pkg/pb"
)

// DataSourceContext 不同数据源之间上下文
type DataSourceContext struct {
}

// DataSource 数据源
type DataSource interface {
	// 读取tick
	GetTick(*DataSourceContext, *pb.Symbol) (*pb.MarketDataSnapshot, error)
	// 取最新一天的tick序列
	GetLastTickSerires(*DataSourceContext, *pb.Symbol) (*pb.TickSeries, error)
	// 读取合约信息
	GetTradingInstrument(*DataSourceContext, *pb.Symbol) (*pb.TradingInstrument, error)
	// 全部合约信息
	TradingInstrumentList(*DataSourceContext, *pb.ReqGetTradingInstrumentList) ([]*pb.TradingInstrument, error)
	// 反向取K线
	RGetKlineSeries(ctx *DataSourceContext, s *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) (*pb.KlineSeries, error)
	// K线
	GetKlineSeries(ctx *DataSourceContext, s *pb.Symbol, period pb.PeriodType, startTime, endTime, lenLimit int64) (*pb.KlineSeries, error)
}

// RealtimeDataSource 实时推送数据源
type RealtimeDataSource interface {
	Subscribe(*pb.ReqSubscribe, chan *pb.MarketDataSnapshot) (*pb.RspSubscribe, error)
	UnSubscribe(*pb.ReqUnSubscribe, chan *pb.MarketDataSnapshot) (*pb.RspUnSubscribe, error)
}

// RPCHandler RPCHandler
type RPCHandler struct {
	realtime RealtimeDataSource
	dsList   []DataSource
}

// MakeRPCHandler 服务
func MakeRPCHandler(dsList []DataSource, realtime RealtimeDataSource) *RPCHandler {
	h := &RPCHandler{realtime: realtime, dsList: dsList}
	return h
}

// Subscribe 订阅行情
func (h *RPCHandler) Subscribe(req *pb.ReqSubscribe, stream pb.DCenter_SubscribeServer) error {
	log.Println("DCenter Subscribe", req)
	ch := make(chan *pb.MarketDataSnapshot, 9999) // fixme: 这个数量需要调整
	h.realtime.Subscribe(req, ch)
	defer h.realtime.UnSubscribe(&pb.ReqUnSubscribe{}, ch)
	for {
		select {
		case md := <-ch:
			err := stream.Send(md)
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

// GetDCenterInfo GetDCenterInfo
func (h *RPCHandler) GetDCenterInfo(ctx context.Context, req *pb.ReqGetDCenterInfo) (*pb.RspGetDCenterInfo, error) {
	res := &pb.RspGetDCenterInfo{}
	return res, nil
}

// GetTradingInstrument GetTradingInstrument
func (h *RPCHandler) GetTradingInstrument(ctx context.Context, req *pb.ReqGetTradingInstrument) (*pb.RspGetTradingInstrument, error) {
	var ret pb.RspGetTradingInstrument
	return &ret, nil
}

// GetTradingInstrumentList ReqGetTradingInstrumentList
func (h *RPCHandler) GetTradingInstrumentList(ctx context.Context, req *pb.ReqGetTradingInstrumentList) (*pb.RspGetTradingInstrumentList, error) {
	var ret pb.RspGetTradingInstrumentList
	return &ret, nil
}

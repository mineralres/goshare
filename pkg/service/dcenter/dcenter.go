package dcenter

import (
	"context"

	pb "github.com/mineralres/goshare/pkg/pb/goshare"
)

// RPCHandler RPCHandler
type RPCHandler struct {
}

// MakeRPCHandler 服务
func MakeRPCHandler() *RPCHandler {
	h := &RPCHandler{}
	return h
}

// Subscribe 订阅行情
func (h *RPCHandler) Subscribe(req *pb.ReqSubscribe, stream pb.DCenter_SubscribeServer) error {
	return nil
}

// CombineSubscribe 综合订阅
func (h *RPCHandler) CombineSubscribe(req *pb.ReqCombineSubscribe, stream pb.DCenter_CombineSubscribeServer) error {
	return nil
}

// GetDCenterInfo GetDCenterInfo
func (h *RPCHandler) GetDCenterInfo(ctx context.Context, req *pb.ReqGetDCenterInfo) (*pb.RspGetDCenterInfo, error) {
	res := &pb.RspGetDCenterInfo{}
	res.CacheSummary = &pb.CacheSummary{}
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

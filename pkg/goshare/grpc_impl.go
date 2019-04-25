package goshare

import (
	context "golang.org/x/net/context"
	"log"
	// "time"
	// "google.golang.org/grpc"

	"errors"
	"github.com/mineralres/goshare/pkg/pb"
)

// GrpcServer GoShareGrpcServer
type GrpcServer struct {
}

// LastTick GetLastTick
func (gs *GrpcServer) LastTick(ctx context.Context, req *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	if req.Exchange > pb.ExchangeType_INVALIDEX {
		var ret pb.MarketDataSnapshot
		return &ret, errors.New("invalid exchange type")
	}
	var ds DataSource
	return ds.GetLastTick(req)
}

// SSEStockOptionList 上证所50ETF期权列表
func (gs *GrpcServer) SSEStockOptionList(ctx context.Context, req *pb.ReqSSEStockOptionList) (*pb.RspSSEStockOptionList, error) {
	var ret pb.RspSSEStockOptionList
	var ds DataSource
	l, err := ds.GetSSEStockOptionList()
	if err == nil {
		for i := range l {
			ret.List = append(ret.List, &l[i])
		}
	}
	return &ret, err
}

// TickStream 行情推送
func (gs *GrpcServer) TickStream(req *pb.Symbol, stream pb.GoShare_TickStreamServer) error {
	var ds DataSource
	for i := 0; i < 10; i++ {
		tick, err := ds.GetLastTick(&pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "601398"})
		if err == nil {
			stream.Send(tick)
		}
	}
	log.Println("stream结束")
	return nil
}

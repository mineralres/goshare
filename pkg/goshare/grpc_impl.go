package goshare

import (
	"errors"
	"log"

	context "golang.org/x/net/context"

	// "time"

	// "errors"
	"github.com/mineralres/goshare/pkg/db"
	"github.com/mineralres/goshare/pkg/pb"
	"google.golang.org/grpc/metadata"
)

// GrpcServer GoShareGrpcServer
type GrpcServer struct {
	cache *db.XCache
}

// 错误定义
var (
	ErrorNeedLogin = errors.New("NeedLogin")
)

// MakeGrpcServer MakeGrpcServer
func MakeGrpcServer() *GrpcServer {
	var ret GrpcServer
	ret.cache = db.MakeXCache()
	return &ret
}

func getUserSession(ctx context.Context) *pb.UserSession {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		md = md
	}
	// v, ok := md["grpcgateway-cookie"]
	return &pb.UserSession{}
}

// LastTick GetLastTick
func (gs *GrpcServer) LastTick(ctx context.Context, req *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	session := getUserSession(ctx)
	if session == nil {
		return &pb.MarketDataSnapshot{}, ErrorNeedLogin
	}
	// do something with metadata
	// 优先查找本地缓存
	ret, err := gs.cache.GetLastTick(req)
	log.Println("lasttick", ret, err, req)
	if err == nil {
		return ret, err
	}
	// 从外部数据源读取
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

// SubscribeMarketData 行情推送
func (gs *GrpcServer) SubscribeMarketData(req *pb.ReqSubscribeMarketData, stream pb.GoShare_SubscribeMarketDataServer) error {
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

// PushTick 推送行情到服务端
func (gs *GrpcServer) PushTick(ctx context.Context, req *pb.MarketDataSnapshot) (*pb.EmptyResponse, error) {
	log.Println("MarketDataSnapshot", req)
	gs.cache.Update(req)
	var ret pb.EmptyResponse
	return &ret, nil
}

// AddTradingRoute 添加经纪准则
func (gs *GrpcServer) AddTradingRoute(ctx context.Context, req *pb.TradingRoute) (*pb.EmptyResponse, error) {
	var ret pb.EmptyResponse
	return &ret, nil
}

// DeleteTradingRoute 删除
func (gs *GrpcServer) DeleteTradingRoute(ctx context.Context, req *pb.ReqDeleteTradingRoute) (*pb.EmptyResponse, error) {
	var ret pb.EmptyResponse
	return &ret, nil
}

// GetTradingRouteList 删除
func (gs *GrpcServer) GetTradingRouteList(ctx context.Context, req *pb.EmptyRequest) (*pb.TradingRouteList, error) {
	var ret pb.TradingRouteList
	return &ret, nil
}

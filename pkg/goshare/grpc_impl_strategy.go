package goshare

import (
	"github.com/mineralres/goshare/pkg/pb"
	context "golang.org/x/net/context"
)

// GetStrategyList 策略列表
func (gs *GrpcServer) GetStrategyList(ctx context.Context, req *pb.EmptyRequest) (*pb.StrategyList, error) {
	var ret pb.StrategyList
	return &ret, nil
}

// GetStrategy 策略列表
func (gs *GrpcServer) GetStrategy(ctx context.Context, req *pb.ReqGetStrategy) (*pb.Strategy, error) {
	var ret pb.Strategy
	return &ret, nil
}

// CreateStrategy 策略列表
func (gs *GrpcServer) CreateStrategy(ctx context.Context, req *pb.Strategy) (*pb.EmptyResponse, error) {
	var ret pb.EmptyResponse
	return &ret, nil
}

// DeleteStrategy 策略列表
func (gs *GrpcServer) DeleteStrategy(ctx context.Context, req *pb.ReqDeleteStrategy) (*pb.EmptyResponse, error) {
	var ret pb.EmptyResponse
	return &ret, nil
}

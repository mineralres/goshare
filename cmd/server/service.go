package main

import (
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/service"
	"github.com/mineralres/goshare/pkg/service/dcenter"
	"github.com/mineralres/goshare/pkg/service/ucenter"
	"google.golang.org/grpc"
)

func runSrv() {
	// 也可以考虑多进程部署
	go runUserManagerSrv()
	go runDCenterSrv()
}

// 用户管理
func runUserManagerSrv() {
	list, err := service.Register("srv.ucenter")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUCenterServer(grpcServer, ucenter.MakeRPCHandler())
	grpcServer.Serve(list)
}

// 行情服务
func runDCenterSrv() {
	list, err := service.Register("srv.dcenter")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDCenterServer(grpcServer, dcenter.MakeRPCHandler())
	grpcServer.Serve(list)
}

package main

import (
	pb "github.com/mineralres/goshare/pkg/pb/goshare"
	"github.com/mineralres/goshare/pkg/service"
	"github.com/mineralres/goshare/pkg/service/dcenter"
	"google.golang.org/grpc"
)

func runSrv(c *config) {
	// 也可以考虑多进程部署
	go runDCenterSrv(c)
}

// 行情服务
func runDCenterSrv(c *config) {
	list, err := service.Register("srv.dcenter")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	// log.Printf("使用goshare官方推送数据源[%s]", c.GSURL.Host)
	pb.RegisterDCenterServer(grpcServer, dcenter.MakeRPCHandler())
	grpcServer.Serve(list)
}

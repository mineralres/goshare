package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mineralres/goshare/pkg/goshare"
	"github.com/mineralres/goshare/pkg/pb"
	"google.golang.org/grpc"
)

func runGrpcService(c xconfig) {
	go func() {
		// 常规服务
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Common.GrpcPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("Common grpc serve on [%d]", c.Common.GrpcPort)
		grpcServer := grpc.NewServer()
		pb.RegisterGoShareServer(grpcServer, goshare.MakeGrpcServer())
		grpcServer.Serve(lis)
	}()
	// 用户服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.User.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("User grpc serve on [%d]", c.User.GrpcPort)
	grpcServer := grpc.NewServer()
	pb.RegisterUserManagerServer(grpcServer, goshare.MakeUserManager())
	grpcServer.Serve(lis)

}

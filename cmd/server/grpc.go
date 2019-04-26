package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mineralres/goshare/pkg/goshare"
	"github.com/mineralres/goshare/pkg/pb"
	"google.golang.org/grpc"
)

func runGrpcService(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Grpc serve on [%d]", port)
	grpcServer := grpc.NewServer()
	pb.RegisterGoShareServer(grpcServer, goshare.MakeGrpcServer())
	grpcServer.Serve(lis)
}

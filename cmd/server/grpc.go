package main

import (
	"fmt"
	"github.com/mineralres/goshare/pkg/goshare"
	"github.com/mineralres/goshare/pkg/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func runGrpcService(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Grpc serve on %d", port)
	grpcServer := grpc.NewServer()
	var impl goshare.GrpcServer
	pb.RegisterGoShareServer(grpcServer, &impl)
	grpcServer.Serve(lis)
}

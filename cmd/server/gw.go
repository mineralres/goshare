package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/mineralres/goshare/pkg/pb"
)

func runGrpcGateway(httpPort int, grpcEndPoint string) error {
	log.Printf("reverseProxy 监听于[%d], grpcEndPoint[%s]", httpPort, grpcEndPoint)
	defer glog.Flush()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterGoShareHandlerFromEndpoint(ctx, mux, grpcEndPoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
}

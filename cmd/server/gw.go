package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/mineralres/goshare/pkg/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func yourMatcher2(headerName string) (mdName string, ok bool) {
	if headerName == "grpc-metadata-set-cookie" {
		return "Set-Cookie", true
	}
	return headerName, true
}

func runGrpcGateway(c xconfig) {
	go func() {
		log.Printf("common gw 监听于[%d]", c.Common.HTTPPort)
		defer glog.Flush()
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false, EmitDefaults: true}),
			runtime.WithOutgoingHeaderMatcher(yourMatcher2))
		opts := []grpc.DialOption{grpc.WithInsecure()}
		gw.RegisterGoShareHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", c.Common.GrpcPort), opts)
		http.ListenAndServe(fmt.Sprintf(":%d", c.Common.HTTPPort), mux)
	}()

	log.Printf("user gw 监听于[%d]", c.User.HTTPPort)
	defer glog.Flush()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false, EmitDefaults: true}),
		runtime.WithOutgoingHeaderMatcher(yourMatcher2))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	gw.RegisterUserManagerHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", c.User.GrpcPort), opts)
	http.ListenAndServe(fmt.Sprintf(":%d", c.User.HTTPPort), mux)

}

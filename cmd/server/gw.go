package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang/glog"
	proto "github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/mineralres/goshare/pkg/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func myFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	// w.Header().Set("X-My-Tracking-Token", resp.Token)
	// resp.Token = ""
	// log.Println("myFilter", w.Header())
	return nil
}

func yourMatcher(headerName string) (mdName string, ok bool) {
	return headerName, true
}

func yourMatcher2(headerName string) (mdName string, ok bool) {
	if headerName == "grpc-metadata-set-cookie" {
		return "Set-Cookie", true
	}
	return headerName, true
}

func runGrpcGateway(httpPort int, grpcEndPoint string) error {
	log.Printf("reverseProxy 监听于[%d], grpcEndPoint[%s]", httpPort, grpcEndPoint)
	defer glog.Flush()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false, EmitDefaults: true}), runtime.WithForwardResponseOption(myFilter),
		runtime.WithIncomingHeaderMatcher(yourMatcher), runtime.WithOutgoingHeaderMatcher(yourMatcher2))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterGoShareHandlerFromEndpoint(ctx, mux, grpcEndPoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
}

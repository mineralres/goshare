package main

import (
	"log"
	"time"

	"context"

	"github.com/micro/go-micro"
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/user"
)

func runSrv(ctx context.Context) {
	runUserManagerSrv(ctx)
}

func runUserManagerSrv(ctx context.Context) {
	service := micro.NewService(
		micro.Name("go.micro.srv.usermanager"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Context(ctx),
	)

	// Register Handlers
	pb.RegisterUserManagerHandler(service.Server(), user.MakeRPCHandler())

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

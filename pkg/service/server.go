package service

import (
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/micro/go-micro/registry"
	"github.com/micro/util/go/lib/addr"
)

// Register 注册一个服务
func Register(name string) (net.Listener, error) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// parse address for host, port
	parts := strings.Split(lis.Addr().String(), ":")
	host := strings.Join(parts[:len(parts)-1], ":")
	port, _ := strconv.Atoi(parts[len(parts)-1])

	ret, err := addr.Extract(host)
	if err != nil {
		return lis, err
	}
	service := &registry.Service{}
	service.Name = name
	service.Version = "latesst"
	node := &registry.Node{
		Id:      uuid.New().String(),
		Address: ret,
		Port:    port,
	}
	service.Nodes = append(service.Nodes, node)
	log.Printf("service [%s] listent at [%s:%d]", name, ret, port)
	return lis, registry.Register(service)
}

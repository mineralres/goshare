package util

import (
	"log"
	"net"
	"strconv"
	"strings"
	"fmt"

	"github.com/google/uuid"
	"github.com/micro/util/go/lib/addr"
	"github.com/micro/go-micro/registry"
	"google.golang.org/grpc"
)

// RegisterService 注册一个服务
func RegisterService(name string) (net.Listener, error) {
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

// GetServiceConn from mdns
func GetServiceConn(name string) (*grpc.ClientConn, error) {
	service, err := registry.GetService(name)
	if err != nil {
		return nil, err
	}
	// log.Println(name, service)
	if len(service) == 0 {
		return nil, fmt.Errorf("service %s not found ", name)
	}
	if len(service[0].Nodes) == 0 {
		return nil, fmt.Errorf("service %s not found ", name)
	}
	node := service[0].Nodes[0]
	addr := fmt.Sprintf("%s:%d", node.Address, node.Port)
	return grpc.Dial(addr, grpc.WithInsecure())
}

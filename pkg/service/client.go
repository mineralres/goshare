package service

import (
	"fmt"

	"github.com/micro/go-micro/registry"
	"google.golang.org/grpc"
)

// GetClientConn from mdns
func GetClientConn(name string) (*grpc.ClientConn, error) {
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

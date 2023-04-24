package service

import (
    "fmt"

	"google.golang.org/grpc"
)


func (s *service) Dial(serviceName, tag string) (*grpc.ClientConn, error) {
    target := fmt.Sprintf("consul://%s/%s?tag=%s", s.consulAddr, serviceName, tag)

	return grpc.Dial(
		target,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithInsecure(),
	)
}

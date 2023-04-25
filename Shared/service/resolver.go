package service

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func (s *service) GRPCListener(server *grpc.Server) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *service) Dial(serviceName, tag string) (*grpc.ClientConn, error) {
	target := fmt.Sprintf("consul://%s/%s?tag=%s", s.consulAddr, serviceName, tag)

    log.Println(target)

	return grpc.Dial(
		target,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithInsecure(),
	)
}

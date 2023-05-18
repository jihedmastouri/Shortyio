package service

import (
	"fmt"
	"log"
	"net"

	_ "github.com/mbobakov/grpc-consul-resolver"

	"google.golang.org/grpc"
)

type DefaultServices string

const (
	Queries DefaultServices = "Queries"
	Search DefaultServices = "Search"
	Web DefaultServices = "Web"
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

func (s *service) Dial(serviceName DefaultServices, tag *[]string) (*grpc.ClientConn, error) {
	consul := s.consulAddr
	t := fmt.Sprintf("consul://%s/%s", consul, serviceName)

	return grpc.Dial(
		t,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
}

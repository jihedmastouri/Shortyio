package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/shorty-io/go-shorty/FlipFlop/handler"
	pb "github.com/shorty-io/go-shorty/FlipFlop/proto"

	"github.com/shorty-io/go-shorty/Shared/consul"
)

func main() {

	consulAddress := os.Getenv("CONSUL_HTTP_ADDR")

	service := consul.NewService()
	service.Register("salah", consulAddress, config.Consul)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCommandsServiceServer(s, &handler.CommandService{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

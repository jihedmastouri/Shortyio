package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/shorty-io/go-shorty/FlipFlop/handler"
	pb "github.com/shorty-io/go-shorty/FlipFlop/proto"
	"google.golang.org/grpc"
)

func main() {

	serverAddress := os.Getenv("SERVER_ADDRESS")

	if serverAddress == "" {
		serverAddress = "localhost:50051"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%v", serverAddress))
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

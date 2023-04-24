package config

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

func Connect() *grpc.ClientConn {
	service := os.Getenv("service")

	if service == "" {
		service = "localhost:50051"
	}

	conn, err := grpc.Dial(service, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	return conn
}

package config

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

func Connect() *grpc.ClientConn {
	flipFlopAddress := os.Getenv("FF_ADDRESS")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	if serverAddress == "" {
		serverAddress = "localhost:50051"
	}

	if flipFlopAddress == "" {
		flipFlopAddress = "localhost:8080"
	}

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	return conn
}

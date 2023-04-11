package config

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

func Connect() *grpc.ClientConn {
	flipFlopAddress := os.Getenv("FF_ADDRESS")

	if flipFlopAddress == "" {
		flipFlopAddress = "localhost:50051"
	}

	conn, err := grpc.Dial(flipFlopAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	return conn
}

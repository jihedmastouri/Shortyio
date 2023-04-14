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

// Dial dials a service through gRPC and returns a new connection.
// func Dial(serviceName, tag string, timeout time.Duration) (grpc.ClientConnInterface, error) {
//     cfg := consul.Config()
//     target := fmt.Sprintf("consul://%s:%s@%s/%s?tag=%s", cfg.User, cfg.Password, cfg.Address, serviceName, tag)
//
// 	return grpc.Dial(
// 		target,
// 		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
// 		grpc.WithInsecure(),
// 	)
// }




	conn, err := grpc.Dial(flipFlopAddress)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	return conn
}

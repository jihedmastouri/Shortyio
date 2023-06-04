package config

import (
	"log"
	"os"

	// f "github.com/shorty-io/go-shorty/flipFlop/proto"
	pb "github.com/shorty-io/go-shorty/Shared/proto"

	"google.golang.org/grpc"
)

type Services struct {
	Queries  pb.QueriesClient
	// FlipFlop f.FlipFlopServiceClient
}


func NewMicroS() *Services{
	queries := os.Getenv("QUERIES")
	// flipflop := os.Getenv("FLIPFLOP")

	connQuery, err := grpc.Dial(queries, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	clientQuery := pb.NewQueriesClient(connQuery)

	// connFlip, err := grpc.Dial(flipflop, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("failed to connect: %v", err)
	// }
	// clientFlip := f.NewFlipFlopServiceClient(connFlip)

	return &Services{
		Queries:  clientQuery,
		// FlipFlop: clientFlip,
	}
}

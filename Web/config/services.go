package config

import (
	"log"
	"os"

	// f "github.com/shorty-io/go-shorty/flipFlop/proto"
	q "github.com/shorty-io/go-shorty/queries/proto"

	"google.golang.org/grpc"
)

type Services struct {
	Queries  q.QueriesClient
	// FlipFlop f.FlipFlopServiceClient
}


func NewMicroS() *Services{
	queries := os.Getenv("QUERIES")
	// flipflop := os.Getenv("FLIPFLOP")

	connQuery, err := grpc.Dial(queries, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	clientQuery := q.NewQueriesClient(connQuery)

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

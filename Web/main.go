package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/FlipFlop/proto"
	"google.golang.org/grpc"
)

func main() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "localhost:50051"
	}

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewCommandsServiceClient(conn)
	req := &pb.CreateCommandRequest{Name: "Mj", Description: "LoL"}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		res, err := client.CreateCommand(context.Background(), req)
		if err != nil {
			return c.String(http.StatusOK, res.GetId())
		}
		return err
	})

	e.Logger.Fatal(e.Start(":80"))
}

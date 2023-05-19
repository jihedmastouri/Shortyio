package main

import (
	"log"
	"net"

	"github.com/shorty-io/go-shorty/flipFlop/handler"
	pb "github.com/shorty-io/go-shorty/flipFlop/proto"
	"github.com/shorty-io/go-shorty/Shared/service"
	"google.golang.org/grpc"
)

func main() {
	srv := service.New("Queries")

	// Not necessary at the moment
	c := service.InitConfig{
		ServiceRegister: service.Consul,
		ConfigProvider:  service.ConsulConfig,
	}
	srv.Init(c)

	srv.Start()

	s := grpc.NewServer()
	pb.RegisterFlipFlopServiceServer(s, &handler.Flip{})

	srv.GRPCListener(s)
}

package main

import (
	"github.com/shorty-io/go-shorty/flipFlop/handler"
	pb "github.com/shorty-io/go-shorty/flipFlop/proto"
	"github.com/shorty-io/go-shorty/Shared/service"
	"google.golang.org/grpc"
)

func main() {
	srv := handler.Srv

	// Not necessary at the moment
	c := service.InitConfig{
		ServiceRegister: service.Consul,
		ConfigProvider:  service.ConsulConfig,
	}
	srv.Init(c)

	srv.Start()

	s := grpc.NewServer()
	pb.RegisterFlipFlopServer(s, &handler.CommandService{})

	srv.GRPCListener(s)
}

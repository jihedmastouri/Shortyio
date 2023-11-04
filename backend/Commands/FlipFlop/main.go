package main

import (
	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"github.com/shorty-io/go-shorty/Shared/service"
	"github.com/shorty-io/go-shorty/Shared/service/namespace"
	"github.com/shorty-io/go-shorty/flipFlop/handler"
	"google.golang.org/grpc"
)

func main() {
	srv := service.New(namespace.FlipFlop)
	handler.NewSrv(srv)

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

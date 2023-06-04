package main

import (
	"time"

	"github.com/shorty-io/go-shorty/Shared/service"
	"github.com/shorty-io/go-shorty/queries/handler"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"google.golang.org/grpc"
)

const TTL = time.Second * 8

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
	pb.RegisterQueriesServer(s, &handler.Queries{})

	srv.GRPCListener(s)
}

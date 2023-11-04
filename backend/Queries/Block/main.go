package main

import (
	"time"

	"google.golang.org/grpc"

	"github.com/shorty-io/go-shorty/queries/db"
	"github.com/shorty-io/go-shorty/queries/handler"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"github.com/shorty-io/go-shorty/Shared/service"
	"github.com/shorty-io/go-shorty/Shared/service/namespace"
)

const TTL = time.Second * 8

func main() {
	srv := service.New(namespace.Queries)

	db.InitConfig(srv)

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

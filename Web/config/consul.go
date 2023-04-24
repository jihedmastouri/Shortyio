package config

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

type Service struct {
    consulClient *api.Client
    config config
}

type config struct {
    user string
    passwd string
    addr string
    name string
}

func NewService()  *Service {
	client, err := api.NewClient(api.DefaultConfig())
    if err != nil {
        log.Fatal(err)
    }

    cfg := newConfig();

    return &Service{
        consulClient: client,
        config: cfg,
    }
}

func ( s *Service) Dial(serviceName, tag string) (grpc.ClientConnInterface, error) {
    cfg := s.config
    target := fmt.Sprintf("consul://%s:%s@%s/%s?tag=%s", cfg.user, cfg.passwd, cfg.addr, serviceName, tag)

	return grpc.Dial(
		target,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithInsecure(),
	)
}

func newConfig() config {

	service := os.Getenv("service")
	user := os.Getenv("user")
	passwd := os.Getenv("passwd")
	addr := os.Getenv("addr")

	if service == "" {
		service = "MyService"
	}

    if addr == "" {
        addr = "0.0.0.0:50051"
    }

    return  config{
        user: user,
        passwd: passwd,
        addr: addr,
        name: service,
    }

}

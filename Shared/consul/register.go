package consul

import (
	"log"
	"time"

	"github.com/hashicorp/consul/api"
)

// Init the Condig

type serviceConfig struct {
	name             string
	address          string
}

func NewService(serviceName, serviceAddress string) serviceConfig {
	return serviceConfig{serviceName, serviceAddress}
}

// Register Service

const (
	TTLInterval                       = time.Second * 15
	TTLRefreshInterval                = time.Second * 10
	TTLDeregisterCriticalServiceAfter = time.Minute
)

func (sc serviceConfig) Register() {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create consul client: %v", err)
	}

	registration := &api.AgentServiceRegistration{
		Name:    sc.name,
		Address: sc.address,
		Port:    50051,
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Failed to register service with consul: %v", err)
	}
}

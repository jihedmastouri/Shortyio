package consul

import (
	"github.com/hashicorp/consul/api"
	"log"
)

// Init the Condig

type serviceConfig struct {
	name             string
	address          string
	serviceDiscovery serviceDiscovery
}


func NewService(serviceName, serviceAddress string, serviceDiscovery serviceDiscovery) *serviceConfig {
	return &serviceConfig{serviceName, serviceAddress, serviceDiscovery}
}

// Register Service: Service Discovery Agnostic

type serviceDiscovery int8

const (
	Consul serviceDiscovery = iota
)

func (sc serviceConfig) Register() {
	switch sc.serviceDiscovery {
	case Consul:
		sc.registerWithConsul()
	}
}

/***
* Registration Mehtods:
***/

// RegisterWithConsul registers the gRPC service with Consul
func (sc serviceConfig) registerWithConsul() {
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

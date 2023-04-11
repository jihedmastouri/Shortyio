package consul

import (
    "github.com/hashicorp/consul/api"
    "google.golang.org/grpc"
    "google.golang.org/grpc/resolver"
)

// RegisterWithConsul registers the gRPC service with Consul
func RegisterWithConsul(serviceName, serviceAddress string) {
    config := api.DefaultConfig()
    client, err := api.NewClient(config)
    if err != nil {
        log.Fatalf("Failed to create consul client: %v", err)
    }

    registration := &api.AgentServiceRegistration{
        Name: serviceName,
        Address: serviceAddress,
        Port: 50051,
    }

    err = client.Agent().ServiceRegister(registration)
    if err != nil {
        log.Fatalf("Failed to register service with consul: %v", err)
    }

    r := &consulResolver{
        client: client,
        serviceName: serviceName,
    }

    resolver.Register(r)
}

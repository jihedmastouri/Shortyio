package service

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	"github.com/simplesurance/grpcconsulresolver/consul"
)

func init() {
	// Register the consul consul at the grpc-go library
	resolver.Register(consul.NewBuilder())
}

func (s *service) GRPCListener(server *grpc.Server) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *service) Dial(serviceName, tag string) (*grpc.ClientConn, error) {

	// consul := os.Getenv("CONSUL_DNS")
	// target := fmt.Sprintf("consul://%s/%s", consul, serviceName)
	target := os.Getenv("target")
	log.Println(target)

	return grpc.Dial(
		target,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithInsecure(),
	)
}

// consulResolver implements the gRPC resolver interface, using Consul for service discovery
// type consulResolver struct {
// 	client *api.Client
// }
//
// func (r *consulResolver) ResolveNow(options resolver.ResolveNowOptions) {}
//
// func (r *consulResolver) Close() {}
//
// func (r *consulResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
// 	// Create a new instance of the consulResolver
// 	// with the provided Consul client
// 	cr := &consulResolver{
// 		client: r.client,
// 	}
//
// 	// Retrieve the service information from Consul
// 	service, _, err := cr.client.Catalog().Service(target.Endpoint(), "", nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Convert the service information to gRPC resolver addresses
// 	addresses := make([]resolver.Address, len(service))
// 	for i, s := range service {
// 		addresses[i] = resolver.Address{
// 			Addr: fmt.Sprintf("%s:%d", s.Address, s.ServicePort),
// 		}
// 	}
//
// 	// Update the client connection with the resolver addresses
// 	cc.UpdateState(resolver.State{
// 		Addresses: addresses,
// 	})
//
// 	return cr, nil
// }
//
// func (r *consulResolver) Scheme() string {
// 	return "consul"
// }

//
// func (r *consulResolver) ResolveNow(resolver.ResolveNowOptions) {}
//
// func (r *consulResolver) Close() {}
//
// func (r *consulResolver) Resolve(target resolver.Target) ([]resolver.Address, error) {
// 	// Retrieve the address and port of the service from Consul
// 	service, _, err := r.client.Catalog().Service(target.Endpoint(), "", nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Convert the service information to gRPC resolver addresses
// 	addresses := make([]resolver.Address, len(service))
// 	for i, s := range service {
// 		addresses[i] = resolver.Address{
// 			Addr: fmt.Sprintf("%s:%d", s.Address, s.ServicePort),
// 		}
// 	}
//
// 	return addresses, nil
// }

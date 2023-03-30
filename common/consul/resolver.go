package consul

import (
	"strconv"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
)


type consulResolver struct {
    client *api.Client
    serviceName string
}

func (r *consulResolver) ResolveNow(resolver.ResolveNowOptions) {}

func (r *consulResolver) Close() {}

func (r *consulResolver) Scheme() string {
    return "consul"
}

func (r *consulResolver) ResolveTarget(target string) (resolver.Resolution, error) {
    serviceEntries, _, err := r.client.Health().Service(r.serviceName, "", true, nil)
    if err != nil {
        return resolver.Resolution{}, err
    }

    var addrs []resolver.Address

    for _, entry := range serviceEntries {
        addr := entry.Service.Address + ":" + strconv.Itoa(entry.Service.Port)
        addrs = append(addrs, resolver.Address{Addr: addr})
    }

    return resolver.Resolution{Addresses: addrs}, nil
}


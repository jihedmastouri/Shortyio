package consul

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func (sc serviceConfig) Dial(serviceName, tag string, timeout time.Duration) {
	dc := dialConfig{serviceName, tag, timeout}

	switch sc.serviceDiscovery {
	case Consul:
		sc.dialConsul(dc)
	}
}

/***
* Dial Mehtods:
***/

type dialConfig struct {
		servicename string
		tag         string
		timeout     time.Duration
}

func (sc serviceConfig) dialConsul( dc dialConfig) (grpc.ClientConnInterface, error) {
	target := fmt.Sprintf("consul://%s/%s?tag=%s", sc.address, dc.servicename, dc.tag)

	return grpc.Dial(
		target,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithInsecure(),
	)
}

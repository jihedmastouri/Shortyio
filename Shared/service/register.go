package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/shorty-io/go-shorty/Shared/service/namespace"
)

const (
	ttl = time.Second * 8 // Time To Live
)

type Service struct {
	agent      *api.Agent
	api        *api.Client
	name       string
	id         string
	consulAddr string
}

type ServiceRegister int8

const (
	Consul ServiceRegister = iota
)

type ConfigProvider int8

const (
	ConsulConfig ConfigProvider = iota
)

type InitConfig struct {
	// Type Of Service Register
	// Default: Consul
	ServiceRegister ServiceRegister

	// Where to get the service Configs from
	// Default: ConsulConfig
	ConfigProvider ConfigProvider
}

// Initialize a new Service
// You must set Consul  Address as Environment variable `CONSUL_HTTP_ADDR`
func New(name namespace.DefaultServices) *Service {
	id := fmt.Sprintf("%s-%s", string(name), uuid.NewString())

	consul := os.Getenv("CONSUL_HTTP_ADDR")
	if consul == "" {
		consul = "localhost:8500"
	}
	log.Println(consul)

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal("New Service:", err)
	}

	agent := client.Agent()

	return &Service{
		agent:      agent,
		api:        client,
		id:         id,
		name:       string(name),
		consulAddr: consul,
	}
}

// For Future Compatibility
func (s *Service) Init(c InitConfig) {}

// Start The Service:
// - Register The service with Consul
// - Listen on a tcp connection for Health updates
func (s *Service) Start() {
	ownAddress, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
	}

	serviceDef := &api.AgentServiceRegistration{
		Name:    s.name,
		ID:      s.id,
		Address: ownAddress,
		Port:    50051,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: ttl.String(),
			TTL:                            ttl.String(),
			CheckID:                        s.id,
			TLSSkipVerify:                  true,
			// HTTP: ":8500",
		},
	}

	if err := s.agent.ServiceRegister(serviceDef); err != nil {
		log.Fatal("Registeration:", err)
	}

	go s.keepAlive()
}

// Update Health status before the Time To Live (TTL) expires
func (s *Service) keepAlive() {
	ticker := time.NewTicker(ttl / 2)
	for {
		if err := s.agent.UpdateTTL(s.id, "online", api.HealthPassing); err != nil {
			log.Fatal(err)
		}
		<-ticker.C
	}
}

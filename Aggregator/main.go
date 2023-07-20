package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/shorty-io/go-shorty/Shared/service"
	"github.com/shorty-io/go-shorty/Shared/service/namespace"
)

var globalConfig map[string]string

func main() {
	srv := service.New(namespace.Aggregator)
	srv.Start()

	globalConfig = initConfig(srv)

	natsUrl := globalConfig["NATS_URL"]
	if natsUrl == "" {
		natsUrl = nats.DefaultURL
	}

	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	_, err = nc.QueueSubscribe("BlockUpdated", "BlockUpdatedQ", BlockUpdated)
	if err != nil {
		log.Fatal(err)
	}

	_, err = nc.QueueSubscribe("BlockDeleted", "BlockDeletedQ", BlockDeleted)
	if err != nil {
		log.Fatal(err)
	}

	if err := nc.Drain(); err != nil {
		log.Fatal(err)
	}

	log.Println("Waiting for messages...")
	select {}
}

func initConfig(srv *service.Service) map[string]string {
	params := []string{
		"MONGO_HOST",
		"MONGO_PASSWORD",
		"MONGO_USER",
		"NATS",
		"POSTGRES_HOST",
		"POSTGRES_PORT",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_DB",
	}

	config := make(map[string]string)
	for _, param := range params {
		value, err := srv.GetKV(param)
		if err != nil {
			log.Fatalf(
				"Failed to retrieve %s from Consul key-value store: %s",
				param,
				err,
			)
		}
		config[param] = value
	}

	return config
}

package main

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"

	"github.com/shorty-io/go-shorty/aggregator/handler"
)

func main() {

	natsUrl := os.Getenv("NATS")
	if natsUrl == "" {
		natsUrl = nats.DefaultURL
	}

	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	_, err = nc.QueueSubscribe("BlockUpdated", "BlockUpdatedQ", handler.BlockUpdated)
	if err != nil {
		log.Fatal(err)
	}

	nc.Flush()
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}
	log.Println("Waiting for messages...")

	select {}
}

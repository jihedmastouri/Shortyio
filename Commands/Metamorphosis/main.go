package main

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
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
    defer nc.Close()
}

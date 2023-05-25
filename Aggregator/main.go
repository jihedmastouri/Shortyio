package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe
	_ , err = nc.QueueSubscribe("BlockUpdated", "BlockUpdatedQ", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	if err != nil {
        log.Fatal(err)
    }

    if err := nc.LastError(); err != nil {
        log.Fatal(err)
    }

}

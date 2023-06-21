package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	defer nc.Close()

	sub, err := nc.QueueSubscribe("BlockUpdate", "BlockUpdatedQ", handler.BlockUpdated)
	if err != nil {
		log.Fatal(err)
	}

	nc.Flush()
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	// Wait for interrupt signal to gracefully exit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down...")
	sub.Unsubscribe() // Unsubscribe from the subject
	time.Sleep(2 * time.Second)
}

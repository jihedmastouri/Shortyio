package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

type Msg struct {
	Id        string
	LangCode  string
	ChangeLog string
}

func BlockUpdated(m *nats.Msg) {
	if m == nil {
		log.Println("Error receiving message")
		m.Nak()
		return
	}

	log.Println("Received a message: ", string(m.Data))

	var msg Msg
	err := json.Unmarshal(m.Data, &msg)
	if err != nil {
		log.Println("failed to unmarshal message:", err)
		m.Nak()
		return
	}

	data, err := AggregateDB(msg)
	if err != nil {
		log.Println("Error aggregating data: ", err)
		m.Nak()
		return
	}

	log.Println("Aggregated data: ", data)

	err = SaveToDB(data)
	if err != nil {
		log.Println("Error saving data: ", err)
		m.Nak()
		return
	}

	log.Println("Data Saved!")
	m.Ack()
}

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

	_, err = nc.QueueSubscribe("BlockUpdated", "BlockUpdatedQ", BlockUpdated)
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

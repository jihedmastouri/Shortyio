package handler

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type Msg struct {
	Id       string
	LangCode string
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

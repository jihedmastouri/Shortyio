package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type Msg struct {
	Id        string
	LangCode  string
	ChangeLog string
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

	_, err = nc.QueueSubscribe("BlockDeleted", "BlockDeletedQ", BlockDeleted)
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

	id, err := uuid.Parse(msg.Id)
	if err != nil {
		log.Println("Error parsing uuid: ", err)
		m.Nak()
		return
	}

	var langs []string
	if msg.LangCode == "" {
		langs, err = getAllLanguages(id)
		if err != nil {
			log.Println("Error getting all languages: ", err)
			m.Nak()
			return
		}
	} else {
		langs = []string{msg.LangCode}
	}

	if len(langs) == 0 {
		log.Println("No languages found")
		langs = []string{"en_US"}
		m.Nak()
		return
	}

	for _, lang := range langs {
		data, err := aggregateDB(id, lang)
		if err != nil {
			log.Println("Error aggregating data: ", err)
			m.Nak()
			return
		}

		err = saveToMongo(data, msg.ChangeLog)
		if err != nil {
			log.Println("Error saving data: ", err)
			m.Nak()
			return
		}
	}

	log.Println("Data Saved!")
	m.Ack()
}

func BlockDeleted(m *nats.Msg) {
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

	id, err := uuid.Parse(msg.Id)
	if err != nil {
		log.Println("Error parsing uuid: ", err)
		m.Nak()
		return
	}

	err = deleteFromMongo(id.String())
	if err != nil {
		log.Println("Error deleting data: ", err)
		m.Nak()
		return
	}

	log.Println("Data Deleted!")
	m.Ack()
}

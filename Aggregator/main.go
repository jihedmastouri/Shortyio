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

	if msg.Id == "" {
		log.Println("Invalid message")
		m.Nak()
		return
	}

	var msgs []Msg

	if msg.LangCode == "" {
		langs, err := getAllLanguages(msg.Id)
		if err != nil {
			log.Println("Error getting all languages: ", err)
			m.Nak()
			return
		}

		for _, lang := range langs {
			msg.LangCode = lang
			msgs = append(msgs, Msg{
				Id:        msg.Id,
				LangCode:  lang,
				ChangeLog: msg.ChangeLog,
			})
		}
	} else {
		msgs = []Msg{msg}
	}

	for _, msg := range msgs {
		data, err := aggregateDB(msg)
		if err != nil {
			log.Println("Error aggregating data: ", err)
			m.Nak()
			return
		}

		log.Println("Aggregated data: ", data)

		err = saveToMongo(data)
		if err != nil {
			log.Println("Error saving data: ", err)
			m.Nak()
			return
		}
	}

	log.Println("Data Saved!")
	m.Ack()
}

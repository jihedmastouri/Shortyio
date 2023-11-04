package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
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

	id, err := uuid.Parse(msg.Id)
	if err != nil {
		log.Println("Error parsing uuid: ", err)
		m.Nak()
		return
	}

	langs := []string{msg.LangCode}

	if msg.LangCode == "" {
		langs, err = getAllLanguages(id)

		if err != nil {
			log.Println("Error getting all languages: ", err)
		}
	}

	if len(langs) == 0 {
		log.Println("No languages found")
		langs = []string{"en_US"}
	}

	var datas [][]byte

	for _, lang := range langs {
		log.Println("Aggregating data for language: ", lang)

		data, err := aggregateDB(id, lang)
		if err != nil {
			log.Println("Error aggregating data: ", err)
			m.Nak()
			return
		}

		datas = append(datas, data)
	}

	for _, data := range datas {
		err = saveToMongo(data, msg.ChangeLog)
		if err != nil {
			// TODO: Publish an event to retry
			log.Println("Error saving data: ", err)
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

package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

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
		}
	} else {
		langs = []string{msg.LangCode}
	}

	if len(langs) == 0 {
		log.Println("No languages found")
		langs = []string{"en_US"}
	}

	for _, lang := range langs {
		log.Println("Aggregating data for language: ", lang)
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

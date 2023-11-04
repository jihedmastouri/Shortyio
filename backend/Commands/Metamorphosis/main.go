package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

var Conn *sql.DB

func BlockUpdated(m *nats.Msg) {
	log.Printf("Received a message: %s\n", string(m.Data))

	var msg MsgU

	err := json.Unmarshal(m.Data, &msg)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	natsUrl := os.Getenv("NATS")
	if natsUrl == "" {
		natsUrl = nats.DefaultURL
	}

	host := os.Getenv("PG_HOST")
	port := "5432"
	user := "postgres"
	password := "root"
	dbname := "shortyio"

	Conn, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	))
	if err != nil {
		log.Fatal(err)
	}
	defer Conn.Close()

	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sub, err := nc.QueueSubscribe("BlockUpdate", "BlockUpdatedQ", BlockUpdated)
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

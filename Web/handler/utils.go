package handler

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/shorty-io/go-shorty/Shared/service/namespace"
	"google.golang.org/grpc"
)

type Dialfn func(namespace.DefaultServices, *[]string) (*grpc.ClientConn, error)

var nc *nats.Conn

func init() {
	natsUrl := os.Getenv("NATS")
	if natsUrl == "" {
		natsUrl = nats.DefaultURL
	}

	var err error
	nc, err = nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}
}

func Cleanup() {
	nc.Flush()
	nc.Close()
}

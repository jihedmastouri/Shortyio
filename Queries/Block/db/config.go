package db

import (
	"context"
	"fmt"
	"log"
	"os"

	// "os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	host := os.Getenv("MONGO_HOST")
	username := os.Getenv("MONGO_USER")
	psswd := os.Getenv("MONGO_PASSWORD")

	connString := fmt.Sprintf(
        "mongodb+srv://%s:%s@%s",
		host,
        username,
        psswd,
    )
	log.Print(connString)

	clientOptions := options.Client().ApplyURI(connString)
	ctx := context.Background()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
        log.Fatal(err)
    }

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("shortyio").Collection("blocks")
}

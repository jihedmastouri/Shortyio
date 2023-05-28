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

func connectMongo() (*mongo.Client, error) {
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
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}

func getCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("shortyio").Collection("blocks")
}

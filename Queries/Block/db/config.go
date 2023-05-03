package db

import (
	"context"
	"fmt"
	"log"
	// "os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {

	pwd := "DWldoNa8losWte27"
	username := "reader"

	// username := os.Getenv("MongoUSR")
	// pwd := os.Getenv("MongoPWD")
	connString := fmt.Sprintf(
        "mongodb+srv://%s:%s@cluster0.ptlgsef.mongodb.net/?w=majority",
        username,
        pwd,
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

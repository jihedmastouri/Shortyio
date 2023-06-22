package db

import (
	"context"
	"fmt"
	"log"

	"github.com/shorty-io/go-shorty/Shared/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = make(map[string]string)

func InitConfig(srv *service.Service) {
	params := []string{
		"MONGO_HOST",
		"MONGO_PASSWORD",
		"MONGO_USER",
	}

	for _, param := range params {
		value, err := srv.GetKV(param)
		if err != nil {
			log.Fatalf(
				"Failed to retrieve %s from Consul key-value store: %s",
				param,
				err,
			)
		}
		config[param] = value
	}
}

func connectMongo(ctx context.Context) (*mongo.Client, error) {

	host := config["MONGO_HOST"]
	username := config["MONGO_USER"]
	psswd := config["MONGO_PASSWORD"]

	connString := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s",
		username,
		psswd,
		host,
	)

	clientOptions := options.Client().ApplyURI(connString)

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

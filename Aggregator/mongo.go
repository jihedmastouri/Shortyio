package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func saveToMongo(data foo, changelog string) error {
	collection, err := connectMongo()
	if err != nil {
		log.Print(err)
		return err
	}

	var bar []any
	if err = json.Unmarshal(data, &bar); err != nil {
		log.Print(err)
		return err
	}

	ctx := context.Background()
	res, err := collection.InsertOne(ctx, bar[0])
	if err != nil {
		log.Println(err)
		return err
	}

	if _, err = collection.UpdateOne(ctx, bson.M{"_id": res.InsertedID},
		bson.M{"$set": bson.M{
			"changelog":  changelog,
			"block_type": "Post",
		}},
	); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func connectMongo() (*mongo.Collection, error) {
	host := "cluster0.ptlgsef.mongodb.net/?retryWrites=true&w=majority"
	username := "reader"
	psswd := "DWldoNa8losWte27"

	connString := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s",
		username,
		psswd,
		host,
	)
	log.Print(connString)

	clientOptions := options.Client().ApplyURI(connString)
	ctx := context.Background()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	collection := client.Database("shortyio").Collection("blocks")
	return collection, nil
}

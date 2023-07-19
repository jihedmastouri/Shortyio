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

func deleteFromMongo(id string) error {
	ctx := context.Background()
	_, err := collection.DeleteMany(ctx, bson.M{"block_id": id})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func saveToMongo(data foo, changelog string) error {
	var bar []any
	if err := json.Unmarshal(data, &bar); err != nil {
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

func connectMongo(config map[string]string) *mongo.Collection {
	host := config["MONGO_HOST"]
	username := config["MONGO_USER"]
	password := config["MONGO_PASSWORD"]

	connString := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s",
		username,
		password,
		host,
	)

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

	return client.Database("shortyio").Collection("blocks")
}

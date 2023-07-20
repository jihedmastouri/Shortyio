package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func deleteFromMongo(id string) error {
	collection, err := connectMongo()
	if err != nil {
		return err
	}

	defer collection.Database().Client().Disconnect(context.Background())

	ctx := context.Background()
	_, err = collection.DeleteMany(ctx, bson.M{"block_id": id})
	if err != nil {
		return fmt.Errorf("Failed to delete data: %w", err)
	}
	return nil
}

func saveToMongo(data []byte, changelog string) error {
	collection, err := connectMongo()
	if err != nil {
		return err
	}

	defer collection.Database().Client().Disconnect(context.Background())

	var bar []any
	if err = json.Unmarshal(data, &bar); err != nil {
		return fmt.Errorf("Failed to unmarshal data: %w", err)
	}

	ctx := context.Background()
	res, err := collection.InsertOne(ctx, bar[0])
	if err != nil {
		return fmt.Errorf("Failed to insert data: %w", err)
	}

	if _, err = collection.UpdateOne(ctx, bson.M{"_id": res.InsertedID},
		bson.M{"$set": bson.M{
			"changelog":  changelog,
			"block_type": "Post",
		}},
	); err != nil {
		return fmt.Errorf("Failed to update data: %w", err)
	}

	return nil
}

func connectMongo() (*mongo.Collection, error) {
	host := globalConfig["MONGO_HOST"]
	username := globalConfig["MONGO_USER"]
	password := globalConfig["MONGO_PASSWORD"]

	connString := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s",
		username,
		password,
		host,
	)

	clientOptions := options.Client().ApplyURI(connString)
	ctx := context.Background()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil || client == nil {
		return nil, fmt.Errorf("Failed to connect to database: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %v", err)
	}

	return client.Database("shortyio").Collection("blocks"), nil
}

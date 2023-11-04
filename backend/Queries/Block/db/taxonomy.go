package db

import (
	"context"
	"errors"
	"log"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"go.mongodb.org/mongo-driver/bson"
)

func getAllTags(ctx context.Context, req *pb.SearchRequest) (*pb.TagList, error) {
	client, err := connectMongo(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("ERROR CONNECTING TO DATABASE")
	}
	defer client.Disconnect(ctx)
	collection := getCollection(client)

	pipeline := []bson.M{
		{"$unwind": "$tags"},
		{"$group": bson.M{
			"_id":     nil,
			"allTags": bson.M{"$addToSet": "$tags"},
		},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return nil, errors.New("ERROR GETTING TAGS")
	}
	defer cursor.Close(ctx)

	var tags struct {
		allTags []string
	}
	if err = cursor.All(ctx, &tags); err != nil {
		log.Println(err)
		return nil, errors.New("ERROR GETTING BLOCKS")
	}

	return &pb.TagList{Tags: tags.allTags}, nil
}

func getAllCategories(ctx context.Context, req *pb.SearchRequest) (*pb.CategoryList, error) {
	client, err := connectMongo(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("ERROR CONNECTING TO DATABASE")
	}
	defer client.Disconnect(ctx)

	collection := getCollection(client)

	pipeline := []bson.M{
		{"$unwind": "$categories"},
		{"$group": bson.M{
			"_id":           nil,
			"allCategories": bson.M{"$addToSet": "$categories"},
		},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return nil, errors.New("ERROR GETTING CATEGORIES")
	}
	defer cursor.Close(ctx)

	var categories struct {
		allCategories []string
	}
	if err = cursor.All(ctx, &categories); err != nil {
		log.Println(err)
		return nil, errors.New("ERROR GETTING BLOCKS")
	}

	return &pb.CategoryList{Categories: categories.allCategories}, nil
}

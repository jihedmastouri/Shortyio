package db

import (
	"context"
	"log"
	"time"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Search(ctx context.Context, req *pb.SearchRequest) (*pb.BlockList, error) {
	client, err := connectMongo(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := getCollection(client)

	query := bson.M{}

	if len(req.Selectors.Tags) > 0 {
		query["tags"] = bson.M{"$in": req.Selectors.Tags}
	}

	if len(req.Selectors.Categories) > 0 {
		query["categories"] = bson.M{"$in": req.Selectors.Categories}
	}

	if len(req.Selectors.Type) > 0 {
		query["type"] = bson.M{
			"$regex":   req.Selectors.Type,
			"$options": "i",
		}
	}

	skip := req.Pagination.PageNum * req.Pagination.PageSize

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(req.Pagination.PageSize))

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var metaList []*pb.BlockMeta
	if err := cursor.All(context.Background(), &metaList); err != nil {
		log.Println(err)
		return nil, err
	}

	totalCount, err := collection.CountDocuments(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.BlockList{
		Metas: metaList,
		Pagination: &pb.Pagination{
			PageNum:  req.Pagination.PageNum + 1,
			PageSize: int32(len(metaList)),
			Total:    &totalCount,
		},
	}, nil
}

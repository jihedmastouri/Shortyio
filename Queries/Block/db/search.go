package db

import (
	"context"
	"errors"
	"log"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"go.mongodb.org/mongo-driver/bson"
)

func Search(ctx context.Context, req *pb.SearchRequest) (*pb.BlockList, error) {
	client, err := connectMongo(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("ERROR CONNECTING TO DATABASE")
	}
	defer client.Disconnect(ctx)
	collection := getCollection(client)

	log.Println("Search Recived:", req.Selectors)

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

	pageSize := 100
	if req.Pagination.PageSize <= 0 && req.Pagination.PageSize > 100 {
		pageSize = int(req.Pagination.PageSize)
	}

	var pagenum int
	if req.Pagination.PageNum > 0 {
		pagenum = int(req.Pagination.PageNum) - 1
	}

	skip := pagenum * pageSize

	pipeline := []bson.M{
		{"$match": query},
		{"$group": bson.M{
			"_id": "block_id", // Specify the field to make distinct
			"doc": bson.M{"$first": "$$ROOT"},
		}},
		{"$replaceRoot": bson.M{"newRoot": "$doc"}},
		{"$skip": int64(skip)},      // Add the $skip stage
		{"$limit": int64(pageSize)}, // Add the $limit stage
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return nil, errors.New("ERROR GETTING BLOCKS")
	}
	defer cursor.Close(ctx)

	var metaList []*pb.BlockMeta

	// Iterate over the cursor
	for cursor.Next(context.Background()) {
		var bm *BlockMeta
		err := cursor.Decode(&bm)
		if err != nil {
			log.Fatal(err)
		}

		// auths := []*pb.Author{}
		// for _, a := range bm.Authors {
		// 	auths = append(auths, &pb.Author{
		// 		Name:  a.Name,
		// 		Id:    a.ID,
		// 		Image: a.Image,
		// 	})
		// }

		metaList = append(metaList, &pb.BlockMeta{
			BlockId:    bm.BlockID,
			Name:       bm.Name,
			Type:       bm.Type,
			Tags:       bm.Tags,
			Categories: bm.Categories,
			// Authors:    auths,
		})
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil, errors.New("ERROR GETTING BLOCKS")
	}

	entries, err := collection.Distinct(ctx, "block_id", query)
	if err != nil {
		log.Println(err)
		return nil, errors.New("ERROR GETTING BLOCKS")
	}
	totalCount := len(entries)
	count := uint32(skip)
	total := uint32(totalCount)

	return &pb.BlockList{
		Metas: metaList,
		Pagination: &pb.Pagination{
			PageNum:  uint32(pagenum) + 1,
			PageSize: uint32(len(metaList)),
			Total:    &total,
			Offset:   &count,
		},
	}, nil
}

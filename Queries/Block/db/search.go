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
	if req.Pagination.PageSize > 0 && req.Pagination.PageSize < uint32(pageSize) {
		pageSize = int(req.Pagination.PageSize)
	}

	var pagenum int
	if req.Pagination.PageNum > 0 {
		pagenum = int(req.Pagination.PageNum) - 1
	}
	skip := pagenum * pageSize

	log.Println(skip, pageSize, pagenum)

	pipeline := []bson.M{
		{"$match": query},
		{"$group": bson.M{
			"_id": "$block_id",
			"doc": bson.M{"$first": "$$ROOT"},
		}},
		{"$replaceRoot": bson.M{"newRoot": "$doc"}},
		{"$skip": int64(skip)},
		{"$limit": int64(pageSize)},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return nil, errors.New("ERROR GETTING BLOCKS")
	}
	defer cursor.Close(ctx)
	log.Println(cursor.RemainingBatchLength())

	var metaList []BlockMeta
	if err = cursor.All(ctx, &metaList); err != nil {
		log.Println(err)
		return nil, errors.New("ERROR GETTING BLOCKS")
	}

	var metaFinal []*pb.BlockMeta
	for _, bm := range metaList {
		log.Println(bm)

		metaFinal = append(metaFinal, &pb.BlockMeta{
			BlockId:    bm.BlockID,
			Name:       bm.Name,
			Type:       bm.Type,
			Tags:       bm.Tags,
			Categories: bm.Categories,
			// Authors:    auths,
		})
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
		Metas: metaFinal,
		Pagination: &pb.Pagination{
			PageNum:  uint32(pagenum) + 1,
			PageSize: uint32(len(metaList)),
			Total:    &total,
			Offset:   &count,
		},
	}, nil
}

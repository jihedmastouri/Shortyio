package db

import (
	"context"

	pb "github.com/shorty-io/go-shorty/queries/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBlockMeta(id, lang string) (*pb.BlockMeta, error) {
	meta := &pb.BlockMeta{}
	query := bson.M{"block_id": id, "lang": lang}
	err := collection.FindOne(context.Background(), query).Decode(meta)
	return meta, err
}

func GetBlock(id, lang string) (*pb.Block, error) {

	query := bson.M{"block_id": id, "lang": lang, "deleted": false}
	cusror := collection.FindOne(context.Background(), query)

	block := &pb.Block{}
	if err := cusror.Decode(block); err != nil {
		return nil, err
	}

	meta := &pb.BlockMeta{}
	if err := cusror.Decode(meta); err != nil {
		return nil, err
	}

	block.Meta = meta
	return block, nil
}

func GetBlockList(selectors pb.Selectors, pagination pb.Pagination) (*pb.BlockList, error) {

	query := bson.M{
		"type":       selectors.GetType(),
		"tags":       selectors.GetTags(),
		"categories": selectors.GetCategories(),
		"authors":    bson.M{"$elemMatch": bson.M{"id": selectors.GetAuthors}},
		"deleted":    false,
	}

	offset := (pagination.GetPageNum - 1) * pagination.GetPageSize()

    pipeline := bson.A{
        bson.M{"$match": query},
        bson.M{"$group": bson.M{"_id": "$" + "block_id", "doc": bson.M{"$first": "$$ROOT"}}},
        bson.M{"$replaceRoot": bson.M{"newRoot": "$doc"}},
        // bson.M{"$count": "count"},
        bson.M{"$skip": offset},
        bson.M{"$limit": pagination.GetPageSize()},
    }

	options := options.Aggregate().SetBatchSize(60)

	cursor, err := collection.Aggregate(context.Background(), pipeline, options)
	if err != nil {
		return nil, err
	}

	var list *pb.BlockList
	if err := cursor.All(context.Background(), list); err != nil {
		return nil, err
	}

    res = *&pb.BlockListResponse{
        blocklist: list,
        pagination: pb.Pagination{
            PageSize: len(list),
            PageNum: pagination.GetPageNum() + 1
            Total: ,
        }
    }

	return results, nil
}

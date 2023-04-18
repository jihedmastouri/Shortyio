package db

import (
	"context"

	pb "github.com/shorty-io/go-shorty/queries/proto"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBlockMeta(id, lang string) (*pb.BlockMeta, error) {
	meta := &pb.BlockMeta{}
	query := bson.M{"block_id": id, "lang": lang}
	err := collection.FindOne(context.Background(), query).Decode(meta)
	return meta, err
}

func GetBlock(id, lang string) (*pb.Block, error) {

	query := bson.M{"block_id": id, "lang": lang}
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

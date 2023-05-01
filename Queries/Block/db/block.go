package db

import (
	"context"

	pb "github.com/shorty-io/go-shorty/queries/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBlockMeta(bq *pb.BlockRequest) (*pb.BlockMeta, error) {

	query, option := buildQuery(bq)
	cusror := collection.FindOne(context.Background(), query, option)

	meta := &pb.BlockMeta{}
	if err := cusror.Decode(meta); err != nil {
		return nil, err
	}

	return meta, nil
}

func GetBlock(bq *pb.BlockRequest) (*pb.Block, error) {

	query, option := buildQuery(bq)
	cusror := collection.FindOne(context.Background(), query, option)

	block := &pb.Block{}
	if err := cusror.Decode(block); err != nil {
		return nil, err
	}

	return block, nil
}

func buildQuery(bq *pb.BlockRequest) (bson.M, *options.FindOneOptions) {

	if bq.GetVersion() != 0 {
		return bson.M{
			"block_id": bq.GetId(),
			"lang":     bq.GetLang(),
			"version":  bq.GetVersion(),
		}, nil
	}

	opts := options.FindOne().SetSort(bson.D{{Key: "version", Value: -1}})
	return bson.M{
		"block_id": bq.GetId(),
		"lang":     bq.GetLang(),
	}, opts
}

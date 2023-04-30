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

func GetBlock(bq pb.BlockRequest) (*pb.Block, error) {

    query, option := buildQuery(bq)
	cusror := collection.FindOne(context.Background(), query, option)

	block := &pb.Block{}
	if err := cusror.Decode(block); err != nil {
		return nil, err
	}

	return block, nil
}

func buildQuery(bq pb.BlockRequest) (bson.M, *options.FindOptions)  {

    if (bq.hasVersion()) {
        return bson.M{
            "block_id": bq.id,
            "lang": bq.lang,
            "version": bq.version,
        }, nil
    }

    opts := options.Find().SetSort(bson.D{{Key: "version", Value: -1}})
    return bson.M{
        "block_id": bq.id,
        "lang": bq.lang,
        "deleted": false,
    }, opts
}



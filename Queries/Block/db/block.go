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

type blockQuery struct {
    id int32
    lang string
    deleted bool
    version int32
}

func buildQuery(bq blockQuery) basn.M, opts {
    opts := options.Find().SetSort(bson.D{{"enrollment", -1}, {"title", 1}})

    if (bq.version != 0) {
        return bson.M{
            "block_id": bq.id,
            "lang": bq.lang,
            "version": version,
        },
    }

    return bson.M{
        "block_id": bq.id,
        "lang": bq.lang,
    }, opts
}



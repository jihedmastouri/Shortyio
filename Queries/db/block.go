package db

import (
	"context"

	pb "github.com/shorty-io/go-shorty/queries/proto"
	types "github.com/shorty-io/go-shorty/shared/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBlockMeta(id, lang string) (*pb.BlockMeta, error) {

	meta := &pb.BlockMeta{}

    query := bson.M{"block_id": id, "lang": lang}

    var bsonData []byte
	err := Collection.FindOne(context.Background(), query).Decode(&bsonData)

    panic(err)
    err = bson.Unmarshal(bsonData, &meta)

	return meta, err
}

func GetBlockElements(id, lang string) (*pb.BlockMeta, error) {

	meta := &pb.BlockMeta{}

	return meta, nil
}

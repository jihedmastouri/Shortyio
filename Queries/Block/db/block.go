package db

import (
	"context"
	"errors"
	"log"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBlockMeta(ctx context.Context, bq *pb.BlockRequest) (*BlockMeta, error) {
	client, err := connectMongo(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Failed To Connect")
	}
	defer client.Disconnect(ctx)
	collection := getCollection(client)

	query, option := buildQuery(bq)
	cusror := collection.FindOne(ctx, query, option)
	meta := new(BlockMeta)
	if err := cusror.Decode(meta); err != nil {
		log.Println(err)
		return nil, errors.New("Failed TO Decode Block")
	}
	return meta, nil
}

func GetBlockContent(ctx context.Context, bq *pb.BlockRequest) (*pb.BlockContent, error) {
	client, err := connectMongo(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := getCollection(client)

	query, option := buildQuery(bq)
	cusror := collection.FindOne(ctx, query, option)
	content := &pb.BlockContent{}
	if err := cusror.Decode(content); err != nil {
		return nil, err
	}
	return content, nil
}

func GetBlock(ctx context.Context, bq *pb.BlockRequest) (*Block, error) {
	client, err := connectMongo(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("CONNECTION ERROR")
	}
	defer client.Disconnect(ctx)
	collection := getCollection(client)

	if bq.GetLang() == "" {
		return nil, errors.New("LANGUAGE NOT FOUND")
	}

	query, option := buildQuery(bq)

	block := new(Block)
	cusror := collection.FindOne(ctx, query, option)
	if err := cusror.Decode(block); err != nil {
		log.Println(err)
		return nil, errors.New("BLOCK NOT FOUND")
	}

	return block, nil
}

func GetLanguages(ctx context.Context, bq *pb.LanguageRequest) (*pb.LanguageList, error) {
	client, err := connectMongo(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := getCollection(client)

	pipeline := bson.A{
		bson.M{"$match": bson.M{"block_id": bq.GetId()}},
		bson.M{"$group": bson.M{"_id": "$lang_code"}},
		bson.M{"$project": bson.M{"_id": 0, "Code": "$_id", "Name": "$lang_name"}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	langList := []*pb.LanguageList_Language{}
	if err = cursor.All(ctx, &langList); err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.LanguageList{
		Langs: langList,
	}, nil
}

func GetVersions(ctx context.Context, bq *pb.VersionsRequest) (*pb.VersionResponse, error) {
	client, err := connectMongo(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := getCollection(client)

	pipeline := bson.A{
		bson.M{"$match": bson.M{"block_id": bq.GetId(), "lang_code": bq.GetLang()}},
		bson.M{"$sort": bson.M{"version": -1, "updated_at": -1}},
		bson.M{"$group": bson.M{"_id": "$version", "changeLog": bson.M{"$first": "$changeLog"}}},
		bson.M{"$project": bson.M{"_id": 0, "version": "$_id", "changeLog": 1}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var vl []struct {
		Version   int32  `bson:"version"`
		ChangeLog string `bson:"changeLog"`
	}

	if err := cursor.All(ctx, &vl); err != nil {
		log.Println(err)
		return nil, err
	}

	var verList []*pb.Ver
	for _, v := range vl {
		verList = append(verList, &pb.Ver{
			Version:   v.Version,
			ChangeLog: &v.ChangeLog,
		})
	}

	return &pb.VersionResponse{
		Versions: verList,
	}, nil
}

func buildQuery(bq *pb.BlockRequest) (bson.M, *options.FindOneOptions) {
	if bq.GetVersion() != 0 {
		return bson.M{
			"block_id":  bq.GetId(),
			"lang_code": bq.GetLang(),
			"version":   bq.GetVersion(),
		}, nil
	}

	opts := options.FindOne().SetSort(bson.D{
		{Key: "version", Value: -1},
		{Key: "updated_at", Value: -1},
	})
	query := bson.M{"block_id": bq.GetId()}
	if bq.GetLang() != "" {
		query["lang_code"] = bq.GetLang()
	}

	return query, opts
}

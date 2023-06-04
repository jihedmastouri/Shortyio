package db

import (
	"context"
	"log"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBlockMeta(bq *pb.BlockRequest) (*pb.BlockMeta, error) {
	client, err := connectMongo()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.Disconnect(context.Background())
	collection := getCollection(client)

	query, option := buildQuery(bq)
	cusror := collection.FindOne(context.Background(), query, option)
	meta := &pb.BlockMeta{}
	if err := cusror.Decode(meta); err != nil {
		return nil, err
	}
	return meta, nil
}

func GetBlockContent(bq *pb.BlockRequest) (*pb.BlockContent, error) {
	client, err := connectMongo()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.Disconnect(context.Background())
	collection := getCollection(client)

	query, option := buildQuery(bq)
	cusror := collection.FindOne(context.Background(), query, option)
	content := &pb.BlockContent{}
	if err := cusror.Decode(content); err != nil {
		return nil, err
	}
	return content, nil
}

func GetBlock(bq *pb.BlockRequest) (*pb.Block, error) {
	client, err := connectMongo()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.Disconnect(context.Background())
	collection := getCollection(client)

	query, option := buildQuery(bq)
	cusror := collection.FindOne(context.Background(), query, option)
	block := &pb.Block{}
	if err := cusror.Decode(block); err != nil {
		log.Println(err)
		return nil, err
	}
	return block, nil
}

func GetLanguages(bq *pb.LanguageRequest) (*pb.LanguageList, error) {
	client, err := connectMongo()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.Disconnect(context.Background())
	collection := getCollection(client)

	pipeline := bson.A{
		bson.M{"$match": bson.M{"block_id": bq.GetId()}},
		bson.M{"$group": bson.M{"_id": "$lang_code"}},
		bson.M{"$project": bson.M{"_id": 0, "lang_code": "$_id"}},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var langList []struct {
		Lang string
	}
	if err = cursor.All(context.Background(), &langList); err != nil {
		log.Println(err)
		return nil, err
	}
	var temp []string
	for _, el := range langList {
		temp = append(temp, el.Lang)
	}

	return &pb.LanguageList{
		Langs: temp,
	}, nil
}

func GetVersions(bq *pb.VersionsRequest) (*pb.VersionResponse, error) {
	client, err := connectMongo()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.Disconnect(context.Background())
	collection := getCollection(client)


	pipeline := bson.A{
		bson.M{"$match": bson.M{"block_id": bq.GetId(), "lang_code": bq.GetLang()}},
		bson.M{"$group": bson.M{"_id": "$version", "changeLog": bson.M{"$first": "$changeLog"}}},
		bson.M{"$project": bson.M{"_id": 0, "version": "$_id", "changeLog": 1}},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var verList []*pb.Ver
	if err := cursor.All(context.Background(), &verList); err != nil {
		log.Println(err)
		return nil, err
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

	opts := options.FindOne().SetSort(bson.D{{Key: "version", Value: -1}})
	return bson.M{
		"block_id":  bq.GetId(),
		"lang_code": bq.GetLang(),
	}, opts
}

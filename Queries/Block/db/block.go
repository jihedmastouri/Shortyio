package db

import (
	"context"
	"log"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Rules struct {
	RuleName          string `bson:"rule_name"`
	Descr             string `bson:"descr"`
	Nested            bool   `bson:"nested"`
	HasLikes          bool   `bson:"has_likes"`
	HasComments       bool   `bson:"has_comments"`
	CommentsNested    bool   `bson:"comments_nested"`
	CommentsHasLike   bool   `bson:"comments_has_likes"`
	CommentsEditable  bool   `bson:"comments_editable"`
	CommentsMaxNested int32  `bson:"comments_max_nested"`
}

type Content struct {
	Elements []struct {
		Media struct {
			Title string `bson:"title"`
			Type  int    `bson:"type"`
			File  string `bson:"file"`
			Alt   string `bson:"alt"`
		} `bson:"media"`
		Text struct {
			Name    string `bson:"name"`
			Type    int    `bson:"type"`
			Content string `bson:"content"`
			Hint    string `bson:"hint"`
		} `bson:"text"`
	} `bson:"elements"`
}

type Block struct {
	pb.BlockMeta
	LangCode string  `bson:"lang_code"`
	Version  int32   `bson:"version"`
	content  Content `bson:"content"`
	Rules    Rules   `bson:"rules"`
}

func GetBlockMeta(ctx context.Context, bq *pb.BlockRequest) (*pb.BlockMeta, error) {
	client, err := connectMongo(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := getCollection(client)

	query, option := buildQuery(bq)
	cusror := collection.FindOne(ctx, query, option)
	meta := &pb.BlockMeta{}
	if err := cusror.Decode(meta); err != nil {
		return nil, err
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

func GetBlock(ctx context.Context, bq *pb.BlockRequest) (*pb.Block, error) {
	client, err := connectMongo(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := getCollection(client)

	query, option := buildQuery(bq)

	// 	&pb.Content_ElementType{
	// 	Element: &pb.Content_ElementType_Text{
	// 		Text: &pb.Content_Textual{
	// 			Name:    "",
	// 			Type:    0,
	// 			Content: "",
	// 			Hint:    "",
	// 		},
	// 	},
	// }
	//

	block := &pb.Block{}

	cusror := collection.FindOne(ctx, query, option)
	if err := cusror.Decode(block); err != nil {
		log.Println(err)
		return nil, err
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
		bson.M{"$group": bson.M{"_id": "$version", "changeLog": bson.M{"$first": "$changeLog"}}},
		bson.M{"$project": bson.M{"_id": 0, "version": "$_id", "changeLog": 1}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var verList []*pb.Ver
	if err := cursor.All(ctx, &verList); err != nil {
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

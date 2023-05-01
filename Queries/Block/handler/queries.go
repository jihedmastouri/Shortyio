package handler

import (
	"context"
	"log"

	db "github.com/shorty-io/go-shorty/queries/db"
	pb "github.com/shorty-io/go-shorty/queries/proto"
)

type Queries struct {
	pb.UnimplementedQueriesServer
}

func (q *Queries) GetBlock(ctx context.Context, rq *pb.BlockRequest) (*pb.Block, error) {
    log.Println("coco")
	return db.GetBlock(rq)
}

func (q *Queries) GetBlockMeta(ctx context.Context, rq *pb.BlockRequest) (*pb.BlockMeta, error) {
	return db.GetBlockMeta(rq)
}

func (q *Queries) GetBlockRules(ctx context.Context, rq *pb.BlockRequest) (*pb.BlockRules, error) {
	block, err := db.GetBlock(rq)
	if err != nil {
		return nil, err
	}
	return block.GetRules(), nil
}

func (q *Queries) GetBlockContent(ctx context.Context, rq *pb.BlockRequest) (*pb.BlockContent, error) {
	block, err := db.GetBlock(rq)
	if err != nil {
		return nil, err
	}

	return &pb.BlockContent{
		BlockId: block.GetBlockId(),
		Content: block.GetContent(),
		Lang:    block.GetLang(),
		Version: block.GetVersion(),
	}, nil
}

func (q *Queries) GetVersions(ctx context.Context, rq *pb.VersionsRequest) (*pb.VersionList, error) {
	return nil, nil
}

func (q *Queries) GetLanguages(ctx context.Context, rq *pb.LanguageRequest) (*pb.LanguageList, error) {
	return nil, nil
}

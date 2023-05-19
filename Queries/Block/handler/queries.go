package handler

import (
	"context"

	db "github.com/shorty-io/go-shorty/queries/db"
	pb "github.com/shorty-io/go-shorty/queries/proto"
)

type Queries struct {
	pb.UnimplementedQueriesServer
}

func (q *Queries) GetBlock(ctx context.Context, rq *pb.BlockRequest) (*pb.Block, error) {
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

// Fix this later
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

func (q *Queries) GetVersions(ctx context.Context, rq *pb.VersionsRequest) (*pb.VersionResponse, error) {
	return db.GetVersions(rq)
}

func (q *Queries) GetLanguages(ctx context.Context, rq *pb.LanguageRequest) (*pb.LanguageList, error) {
	return db.GetLanguages(rq)
}

package handler

import (
	"context"

	"github.com/shorty-io/go-shorty/queries/db"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

type Queries struct {
	pb.UnimplementedQueriesServer
}

func (q *Queries) GetBlock(ctx context.Context, rq *pb.BlockRequest) (*pb.Block, error) {
	return db.GetBlock(ctx, rq)
}

func (q *Queries) GetBlockMeta(ctx context.Context, rq *pb.BlockRequest) (*pb.BlockMeta, error) {
	return db.GetBlockMeta(ctx, rq)
}

func (q *Queries) GetBlockRules(ctx context.Context, rq *pb.BlockRequest) (*pb.BlockRules, error) {
	block, err := db.GetBlock(ctx, rq)
	if err != nil {
		return nil, err
	}
	return block.GetRules(), nil
}

// Fix this later
func (q *Queries) GetBlockContent(ctx context.Context, rq *pb.BlockRequest) (*pb.BlockContent, error) {
	block, err := db.GetBlock(ctx, rq)
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
	return db.GetVersions(ctx, rq)
}

func (q *Queries) GetLanguages(ctx context.Context, rq *pb.LanguageRequest) (*pb.LanguageList, error) {
	return db.GetLanguages(ctx, rq)
}

func (q *Queries) Search(ctx context.Context, rq *pb.SearchRequest) (*pb.BlockList, error) {
	return db.Search(ctx, rq)
}

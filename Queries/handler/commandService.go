package handler

import (
	"context"
	"fmt"

	pb "github.com/shorty-io/go-shorty/queries/proto"
	db "github.com/shorty-io/go-shorty/queries/db"
)

type Queries struct {
	pb.UnimplementedQueriesServer
}

func (q *Queries) GetBlock(ctx context.Context, rq *pb.BlockID) (*pb.Block, error) {

    blockMeta := pb.BlockMeta{}
    block := pb.Block{blockMeta}
	return block, nil
}

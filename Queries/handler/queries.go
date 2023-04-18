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

    block, err := db.GetBlock(rq.GetId() ,rq.GetLang())
    if err != nil {
        return nil, err
    }
	return block, nil
}

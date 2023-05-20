package handler

import (
	"context"
	"database/sql"

	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/flipFlop/proto"
)

type CommandService struct {
	pb.UnimplementedFlipFlopServer
}

func (c *CommandService) CreateBlock(rq *pb.CreateRequest) (*pb.ActionResponse, error) {
	ctx := context.Background()

	conn, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		return err
	}

	q := db.New(conn)


	q.AddBlock(ctx, db.AddBlockParams{
		HasLikes:     sql.NullBool{},
		HasComments:  sql.NullBool{},
		BlockType:    uuid.NullUUID{},
		CommentsType: rq.Meta.,
	})

	return &pb.ActionResponse{
		IsSuceess: false,
		Id:        "",
		Message:   "",
	}, nil
}

func (c *CommandService) CreateBlockLang(*pb.CreateLangRequest) (*pb.ActionResponse, error) {
	return &pb.ActionResponse{
		IsSuceess: false,
		Id:        "",
		Message:   "",
	}, nil
}
func (c *CommandService) DeleteBlock(*pb.DeleteRequest) (*pb.ActionResponse, error) {
	return &pb.ActionResponse{
		IsSuceess: false,
		Id:        "",
		Message:   "",
	}, nil
}

func (c *CommandService) DeleteBlockLang(*pb.DeleteLangRequest) (*pb.ActionResponse, error) {
	return &pb.ActionResponse{
		IsSuceess: false,
		Id:        "",
		Message:   "",
	}, nil
}

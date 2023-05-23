package handler

import (
	"context"
	"database/sql"
	"log"

	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/flipFlop/proto"
)

func (c *CommandService) CreateBlockRule(rq *pb.BlockRules) (*pb.ActionResponse, error) {
	var conn *sql.DB
	if err := newConn(conn); err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Message:   "Failed to connect to database",
		}, nil
	}
	defer conn.Close()
	q := db.New(conn)

	if rq.GetRules() == nil {
		log.Print("No Rules Provided")
		return &pb.ActionResponse{
			IsSuceess: false,
			Message:   "No Rules Provided",
		}, nil
	}

	ctx := context.Background()
	params := db.AddBlockRuleParams{
		Name: "",
		Nested: sql.NullBool{
			Bool:  rq.GetRules().GetNested(),
			Valid: true,
		},
		HasLikes: sql.NullBool{
			Bool:  rq.GetRules().GetHasLikes(),
			Valid: true,
		},
		HasComments: sql.NullBool{
			Bool:  rq.GetRules().GetHasComments(),
			Valid: true,
		},
		CommentsMaxNest: sql.NullInt16{
			Int16: int16(rq.GetRules().GetCommentsMaxNested()),
			Valid: true,
		},
		CommentsHasLikes: sql.NullBool{
			Bool:  rq.GetRules().GetCommentsHasLikes(),
			Valid: true,
		},
		CommentEditable: sql.NullBool{
			Bool:  rq.GetRules().GetCommentsEditable(),
			Valid: true,
		},
	}

	id, err := q.AddBlockRule(ctx, params)
	if err != nil {
		log.Print("Failed to parse author UUID:", err)
		return nil, err
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id,
		Message:   "Great Success",
	}, nil
}

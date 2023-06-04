package handler

import (
	"context"
	"database/sql"
	"log"

	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/Shared/proto"


	_ "github.com/lib/pq"
)

func (c *CommandService) CreateBlockRule(ctx context.Context, rq *pb.BlockRules) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
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

	params := db.CreateBlockRuleParams{
		Name: rq.GetRules().RuleName,
		Nested: sql.NullBool{
			Bool:  rq.GetRules().Nested,
			Valid: true,
		},
		HasLikes: sql.NullBool{
			Bool:  rq.GetRules().HasLikes,
			Valid: true,
		},
		HasComments: sql.NullBool{
			Bool:  rq.GetRules().HasComments,
			Valid: true,
		},
		CommentsMaxNest: sql.NullInt16{
			Int16: int16(rq.GetRules().CommentsMaxNested),
			Valid: true,
		},
		CommentsHasLikes: sql.NullBool{
			Bool:  rq.GetRules().CommentsHasLikes,
			Valid: true,
		},
		CommentEditable: sql.NullBool{
			Bool:  rq.GetRules().CommentsEditable,
			Valid: true,
		},
	}

	id, err := q.CreateBlockRule(ctx, params)
	if err != nil {
		return nil, err
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id,
		Message:   "Great Success",
	}, nil
}

// TODO: CHANGE THIS TO UPDATE INSTEAD OF DELETE+CREATE
func (*CommandService) UpdateBlockRule(ctx context.Context, rq *pb.BlockRules) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Message:   "Failed to connect to database",
		}, nil
	}
	defer conn.Close()
	q := db.New(conn)

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Message:   "Failed to connect to database",
		}, nil
	}
	defer tx.Rollback()

	if rq.GetRules() == nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Message:   "No Rules Provided",
		}, nil
	}

	if err = q.DeleteBlockRule(ctx, rq.GetRules().RuleName); err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Message:   "Failed to Delete Rule",
		}, nil
	}

	params := db.CreateBlockRuleParams{
		Name: rq.GetRules().RuleName,
		Nested: sql.NullBool{
			Bool:  rq.GetRules().Nested,
			Valid: true,
		},
		HasLikes: sql.NullBool{
			Bool:  rq.GetRules().HasLikes,
			Valid: true,
		},
		HasComments: sql.NullBool{
			Bool:  rq.GetRules().HasComments,
			Valid: true,
		},
		CommentsMaxNest: sql.NullInt16{
			Int16: int16(rq.GetRules().CommentsMaxNested),
			Valid: true,
		},
		CommentsHasLikes: sql.NullBool{
			Bool:  rq.GetRules().CommentsHasLikes,
			Valid: true,
		},
		CommentEditable: sql.NullBool{
			Bool:  rq.GetRules().CommentsEditable,
			Valid: true,
		},
	}

	id, err := q.CreateBlockRule(ctx, params)
	if err != nil {
		return nil, err
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id,
		Message:   "Great Success",
	}, nil
}

func (*CommandService) DeleteBlockRule(ctx context.Context, rq *pb.BlockRules) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Message:   "Failed to connect to database",
		}, nil
	}
	defer conn.Close()
	q := db.New(conn)

	var ruleName string
	if rq.GetRules() == nil {
		ruleName = rq.GetRuleName()
	} else {
		ruleName = rq.GetRules().RuleName
	}

	if err = q.DeleteBlockRule(ctx, ruleName); err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Message:   "Failed to Delete Rule",
		}, nil
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Message:   "Great Success",
	}, nil
}



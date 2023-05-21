package handler

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
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
		log.Print("Database Connection Failed", err)
		return nil, err
	}

	q := db.New(conn)

	author, err := uuid.Parse(rq.GetAuthor())
	if err != nil {
		log.Print("Failed to parse author UUID:", err)
		return nil, err
	}

	rules, name_rule := prepareRules(q, rq.GetRules())

	params := db.AddBlockParams{
		Author:           author,
		Name:             rq.Meta.GetName(),
		Nested:           rules.GetNested(),
		HasLikes:         rules.GetHasLikes(),
		HasComments:      rules.GetHasComments(),
		CommentsMaxNest:  int16(rules.GetCommentsMaxNested()),
		CommentsHasLikes: rules.GetCommentsHasLikes(),
		CommentEditable:  rules.GetCommentsEditable(),
		RulesName:        sql.NullString{String: name_rule, Valid: true},
		Type:             0,
	}

	id, err := q.AddBlock(ctx, params)
	if err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   err.Error(),
		}, nil
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "",
	}, nil
}


func (c *CommandService) UpdateBlock(rq *pb.CreateRequest) (*pb.ActionResponse, error) {
	ctx := context.Background()

	conn, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Print("Database Connection Failed", err)
		return nil, err
	}

	q := db.New(conn)

	if err != nil {
		log.Print("Failed to parse author UUID:", err)
		return nil, err
	}

	rules, name_rule := prepareRules(q, rq.GetRules())

	params := db.UpdateBlockParams{
		Id:				  rq.GetMeta().BlockId,
		Name:             rq.GetMeta().Name,
		RulesName:        sql.NullString{String: name_rule, Valid: true},
		Nested:           rules.GetNested(),
		HasLikes:         rules.GetHasLikes(),
		HasComments:      rules.GetHasComments(),
		CommentsMaxNest:  int16(rules.GetCommentsMaxNested()),
		CommentsHasLikes: rules.GetCommentsHasLikes(),
		CommentEditable:  rules.GetCommentsEditable(),
	}

	err := q.UpdateBlock(ctx, params)
	if err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   err.Error(),
		}, nil
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
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

func (c *CommandService) CreateBlockLang(*pb.CreateLangRequest) (*pb.ActionResponse, error) {
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

func prepareRules(q *db.Queries, br *pb.BlockRules) (pb.BlockRules_Rules, string) {

	if br.GetRules() != nil {
		return *br.GetRules(), "custom"
	}

	ctx := context.Background()
	rules, err := q.GetBlockRulesByName(ctx, br.GetRuleName())
	if err != nil {
		panic(err)
	}

	return pb.BlockRules_Rules{
		Nested:            rules.Nested.Bool,
		HasLikes:          rules.HasLikes.Bool,
		HasComments:       rules.HasComments.Bool,
		CommentsNested:    rules.Nested.Bool,
		CommentsHasLikes:  rules.CommentsHasLikes.Bool,
		CommentsEditable:  rules.CommentEditable.Bool,
		CommentsMaxNested: int32(rules.CommentsMaxNest.Int16),
	}, br.GetRuleName()
}

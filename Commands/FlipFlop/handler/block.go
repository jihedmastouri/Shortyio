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
	conn, err := newConn()
	if err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to connect to database",
		}, nil
	}

	defer conn.Close()
	q := db.New(conn)

	author, err := uuid.Parse(rq.GetAuthor())
	if err != nil {
		log.Print("Failed to parse author UUID:", err)
		return nil, err
	}

	rules, name_rule := getBlockRules(q, rq.GetRules())

	params := db.CreateBlockParams{
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

	ctx := context.Background()

	id, err := q.CreateBlock(ctx, params)
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
	conn, err := newConn()
	if err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to connect to database",
		}, nil
	}

	defer conn.Close()
	q := db.New(conn)


	id, err := uuid.Parse(rq.GetMeta().BlockId)
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, err
	}

	rules, name_rule := getBlockRules(q, rq.GetRules())

	params := db.UpdateBlockParams{
		ID:               id,
		Name:             rq.GetMeta().Name,
		RulesName:        sql.NullString{String: name_rule, Valid: true},
		Nested:           rules.GetNested(),
		HasLikes:         rules.GetHasLikes(),
		HasComments:      rules.GetHasComments(),
		CommentsMaxNest:  int16(rules.GetCommentsMaxNested()),
		CommentsHasLikes: rules.GetCommentsHasLikes(),
		CommentEditable:  rules.GetCommentsEditable(),
	}

	ctx := context.Background()
	if err = q.UpdateBlock(ctx, params); err != nil {
		log.Print("Failed to delete block:", err)
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   err.Error(),
		}, nil
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Updated successfully",
	}, nil
}

func (c *CommandService) DeleteBlock(rq *pb.DeleteRequest) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to connect to database",
		}, nil
	}

	defer conn.Close()
	q := db.New(conn)


	ctx := context.Background()
	id, err := uuid.Parse(rq.GetId())
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, err
	}

	if err = q.DeleteBlock(ctx, id); err != nil {
		log.Print("Failed to delete block:", err)
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        id.String(),
			Message:   "Failed to delete block",
		}, nil
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Deleted successfully",
	}, nil
}

func (c *CommandService) CreateBlockLang(rq *pb.CreateLangRequest) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to connect to database",
		}, nil
	}

	defer conn.Close()
	q := db.New(conn)


	blockid, err := uuid.Parse(rq.BlockId)
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, err
	}

	ctx := context.Background()
	params := db.CreateLangParams{
		LangName: rq.Id,
		LangCode: rq.LangName,
		BlockID:  blockid,
	}

	id, err := q.CreateLang(ctx, params)
	if err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   err.Error(),
		}, nil
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        string(id),
		Message:   "Deleted successfully",
	}, nil
}

func (c *CommandService) DeleteBlockLang(rq *pb.DeleteLangRequest) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to connect to database",
		}, nil
	}

	defer conn.Close()
	q := db.New(conn)


	id, err := uuid.Parse(rq.GetId())
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, err
	}

	ctx := context.Background()
	params := db.DeleteBlockLangParams{
		BlockID:  id,
		LangCode: rq.GetLangName(),
	}

	if err = q.DeleteBlockLang(ctx, params); err != nil {
		log.Print("Failed to delete block lang:", err)
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        id.String(),
			Message:   "Failed to delete block",
		}, nil
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Deleted successfully",
	}, nil
}

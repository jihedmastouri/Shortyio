package handler

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/google/uuid"
	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/Shared/proto"

	_ "github.com/lib/pq"
)

type CommandService struct {
	pb.UnimplementedFlipFlopServer
}

func (c *CommandService) CreateBlock(ctx context.Context, rq *pb.CreateRequest) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.New("FAILED TO CONNECT TO DATABASE")
	}

	defer conn.Close()
	q := db.New(conn)

	author, err := uuid.Parse(rq.GetAuthor())
	if err != nil {
		log.Print("Failed to parse author UUID:", err)
		return nil, errors.New("FAILED TO PARSE AUTHOR ID")
	}

	rules := getBlockRules(q, rq.GetRules())

	blockType, err := q.GetTypeByName(ctx, rq.BlockType)
	if err != nil {
		log.Print("Failed to get block type:", err)
		return nil, errors.New("FAILED TO GET BLOCK TYPE")
	}

	params := db.CreateBlockParams{
		Author:           author,
		Name:             rq.Name,
		Nested:           rules.GetNested(),
		HasLikes:         rules.GetHasLikes(),
		HasComments:      rules.GetHasComments(),
		CommentsMaxNest:  int16(rules.GetCommentsMaxNested()),
		CommentsHasLikes: rules.GetCommentsHasLikes(),
		CommentEditable:  rules.GetCommentsEditable(),
		RulesName:        sql.NullString{String: rules.RuleName, Valid: true},
		Type:             blockType,
	}

	id, err := q.CreateBlock(ctx, params)
	if err != nil {
		log.Print("Failed to create block:", err)
		return nil, errors.New("FAILED TO CREATE BLOCK")
	}

	publishEvent(Msg{
		Id:        id.String(),
		LangCode:  "en",
		ChangeLog: "Created Block",
	})

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "",
	}, nil
}

func (c *CommandService) UpdateBlock(ctx context.Context, rq *pb.CreateRequest) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.New("FAILED TO CONNECT TO DATABASE")
	}

	defer conn.Close()
	q := db.New(conn)

	id, err := uuid.Parse(rq.GetId())
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, errors.New("FAILED TO PARSE BLOCK ID")
	}

	rules := getBlockRules(q, rq.GetRules())

	params := db.UpdateBlockParams{
		ID:               id,
		Name:             rq.Name,
		RulesName:        sql.NullString{String: rules.RuleName, Valid: true},
		Nested:           rules.GetNested(),
		HasLikes:         rules.GetHasLikes(),
		HasComments:      rules.GetHasComments(),
		CommentsMaxNest:  int16(rules.GetCommentsMaxNested()),
		CommentsHasLikes: rules.GetCommentsHasLikes(),
		CommentEditable:  rules.GetCommentsEditable(),
	}

	if err = q.UpdateBlock(ctx, params); err != nil {
		log.Print("Failed to delete block:", err)
		return nil, errors.New("FAILED TO UPDATE BLOCK")
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Updated successfully",
	}, nil
}

func (c *CommandService) DeleteBlock(ctx context.Context, rq *pb.DeleteRequest) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.New("FAILED TO CONNECT TO DATABASE")
	}

	defer conn.Close()
	q := db.New(conn)

	id, err := uuid.Parse(rq.GetId())
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, errors.New("FAILED TO PARSE BLOCK ID")
	}

	if err = q.DeleteBlock(ctx, id); err != nil {
		log.Print("Failed to delete block:", err)
		return nil, errors.New("FAILED TO DELETE BLOCK")
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Deleted successfully",
	}, nil
}

func (c *CommandService) CreateBlockLang(ctx context.Context, rq *pb.CreateLangRequest) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.New("FAILED TO CONNECT TO DATABASE")
	}

	defer conn.Close()
	q := db.New(conn)

	blockid, err := uuid.Parse(rq.BlockId)
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, errors.New("FAILED TO PARSE BLOCK ID")
	}

	params := db.CreateLangParams{
		LangName: rq.LangName,
		LangCode: rq.LangCode,
		BlockID:  blockid,
	}

	id, err := q.CreateLang(ctx, params)
	if err != nil {
		log.Print("Failed to create block:", err)
		return nil, errors.New("FAILED TO CREATE BLOCK")
	}

	publishEvent(Msg{
		Id:        rq.BlockId,
		LangCode:  rq.LangCode,
		ChangeLog: "Created BLOCKLANG",
	})

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        strconv.Itoa(int(id)),
		Message:   "Created successfully",
	}, nil
}

func (c *CommandService) DeleteBlockLang(ctx context.Context, rq *pb.DeleteLangRequest) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.New("FAILED TO CONNECT TO DATABASE")
	}

	defer conn.Close()
	q := db.New(conn)

	id, err := uuid.Parse(rq.GetId())
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, errors.New("FAILED TO PARSE BLOCK ID")
	}

	params := db.DeleteBlockLangParams{
		BlockID:  id,
		LangCode: rq.LangCode,
	}

	if err = q.DeleteBlockLang(ctx, params); err != nil {
		log.Print("Failed to delete block lang:", err)
		return nil, errors.New("FAILED TO DELETE BLOCK LANG")
	}

	publishEvent(Msg{
		Id:        id.String(),
		LangCode:  rq.LangCode,
		ChangeLog: "DELETED BLOCKLANG",
	})

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Deleted successfully",
	}, nil
}

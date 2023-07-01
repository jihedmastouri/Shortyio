package handler

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

func (c *CommandService) UpdateBlockRules(ctx context.Context, rq *pb.BlockRulesRq) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.New("FAILED TO CONNECT TO DATABASE")
	}

	defer conn.Close()
	q := db.New(conn)

	id, err := uuid.Parse(rq.GetBlockId())
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, errors.New("FAILED TO PARSE BLOCK ID")
	}

	if err = q.UpdateBlockRules(ctx, db.UpdateBlockRulesParams{
		ID:               id,
		RulesName:        sql.NullString{String: rq.GetRules().GetRuleName(), Valid: true},
		Nested:           rq.GetRules().GetNested(),
		HasLikes:         rq.GetRules().GetHasLikes(),
		HasComments:      rq.GetRules().GetHasComments(),
		CommentsMaxNest:  int16(rq.GetRules().GetCommentsMaxNested()),
		CommentsHasLikes: rq.GetRules().GetCommentsHasLikes(),
		CommentEditable:  rq.GetRules().GetCommentsEditable(),
	}); err != nil {
		log.Print("Failed to update block:", err)
		return nil, errors.New("FAILED TO UPDATE BLOCK")
	}

	publishEvent(Msg{
		Id:        id.String(),
		LangCode:  "",
		ChangeLog: "Block Rules Updated",
	})

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Updated successfully",
	}, nil
}

func (c *CommandService) UpdateBlockMeta(ctx context.Context, rq *pb.BlockMeta) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.New("FAILED TO CONNECT TO DATABASE")
	}

	defer conn.Close()
	q := db.New(conn)

	id, err := uuid.Parse(rq.GetBlockId())
	if err != nil {
		log.Print("Failed to parse Block UUID:", err)
		return nil, errors.New("FAILED TO PARSE BLOCK ID")
	}

	if err = q.UpdateBlockMeta(ctx, db.UpdateBlockMetaParams{
		ID:          id,
		Name:        rq.GetName(),
		Description: sql.NullString{String: rq.GetDescription(), Valid: true},
	}); err != nil {
		log.Print("Failed to update block:", err)
		return nil, errors.New("FAILED TO UPDATE BLOCK")
	}

	for _, tag := range rq.GetTags() {
		if err = q.AddTagToBlock(ctx, db.AddTagToBlockParams{
			BlockID: id,
			Name:    tag,
		}); err != nil {
			log.Print("Failed to add tag to block:", err)
		}
	}

	for _, category := range rq.GetCategories() {
		if err = q.AddCategToBlock(ctx, db.AddCategToBlockParams{
			BlockID: id,
			Name:    category,
		}); err != nil {
			log.Print("Failed to add category to block:", err)
		}
	}

	publishEvent(Msg{
		Id:        id.String(),
		LangCode:  "",
		ChangeLog: "Block Info Updated",
	})

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Updated successfully",
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
		Description:      sql.NullString{String: rq.Description, Valid: true},
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

	publishEvent(Msg{
		Id:        id.String(),
		LangCode:  "",
		ChangeLog: "Block Updated",
	})

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id.String(),
		Message:   "Updated successfully",
	}, nil
}

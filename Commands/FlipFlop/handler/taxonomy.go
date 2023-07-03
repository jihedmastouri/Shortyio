package handler

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/Shared/proto"

	_ "github.com/lib/pq"
)

// CreateTag creates a new tag
func (s *CommandService) CreateTag(ctx context.Context, in *pb.CreateTaxonomy) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.New("FAILED TO CONNECT TO DATABASE")
	}

	defer conn.Close()
	q := db.New(conn)

	if in.Name == "" {
		return nil, errors.New("NAME CANNOT BE EMPTY")
	}

	params := db.CreateTagParams{
		Name: in.GetName(),
		Descr: sql.NullString{
			String: in.Descr,
			Valid:  true,
		},
	}

	id, err := q.CreateTag(ctx, params)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.New("FAILED TO CREATE TAG")
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id,
		Message:   "Tag created",
	}, nil
}

// CreateCategory creates a new category
func (s *CommandService) CreateCategory(ctx context.Context, in *pb.CreateTaxonomy) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.New("FAILED TO CONNECT TO DATABASE")
	}

	defer conn.Close()
	q := db.New(conn)

	if in.Name == "" {
		return nil, errors.New("NAME CANNOT BE EMPTY")
	}

	params := db.CreateCategParams{
		Name: in.Name,
		Descr: sql.NullString{
			String: in.Descr,
			Valid:  false,
		},
	}

	id, err := q.CreateCateg(ctx, params)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return nil, errors.New("FAILED TO CREATE CATEGORY")
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id,
		Message:   "Category created",
	}, nil
}

// DeleteTag deletes a tag
func (s *CommandService) DeleteTag(ctx context.Context, rq *pb.DeleteTaxonomy) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.New("FAILED TO CONNECT TO DATABASE")
	}

	defer conn.Close()
	q := db.New(conn)

	id, err := q.DeleteTag(ctx, rq.Name)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return nil, errors.New("FAILED TO DELETE TAG")
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        string(id),
		Message:   "Tag deleted",
	}, nil
}

// DeleteCategory deletes a category
func (s *CommandService) DeleteCategory(ctx context.Context, rq *pb.DeleteTaxonomy) (*pb.ActionResponse, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.New("FAILED TO CONNECT TO DATABASE")
	}

	defer conn.Close()
	q := db.New(conn)

	id, err := q.DeleteCateg(ctx, rq.Name)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return nil, errors.New("FAILED TO DELETE CATEGORY")
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        string(id),
		Message:   "Category deleted",
	}, nil
}

func (s *CommandService) JoinTag(ctx context.Context, rq *pb.JoinTaxonomy) (*pb.ActionResponse, error) {
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

	for _, tag := range rq.Names {
		if tag == "" {
			continue
		}

		row, err := q.GetTagByName(ctx, tag)
		if err != nil {
			log.Print("Failed to get tag by name:", err)
		}

		if row.Name == "" || err == sql.ErrNoRows {
			_, err := q.CreateTag(ctx, db.CreateTagParams{
				Name: tag,
				Descr: sql.NullString{
					String: "",
					Valid:  false,
				},
			})
			if err != nil {
				return nil, errors.New("FAILED TO CREATE New TAG")
			}
		}

		params := db.AddTagToBlockParams{
			BlockID: blockid,
			Name:    tag,
		}

		err = q.AddTagToBlock(ctx, params)

		if err != nil {
			log.Print("Failed to add tag to block:", err)
			return nil, errors.New("FAILED TO JOIN TAG")
		}
	}

	publishEvent(Msg{
		Id:        blockid.String(),
		LangCode:  "",
		ChangeLog: "Added Categs",
	}, "BlockUpdated")

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        "",
		Message:   "Tag added to block",
	}, nil
}

func (s *CommandService) JoinCategory(ctx context.Context, rq *pb.JoinTaxonomy) (*pb.ActionResponse, error) {
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

	for _, categ := range rq.Names {

		row, err := q.GetCategoryByName(ctx, categ)
		if err != nil {
			log.Print("Failed to get categ by name:", err)
		}

		if row.Name == "" || err == sql.ErrNoRows {
			_, err := q.CreateCateg(ctx, db.CreateCategParams{
				Name: categ,
				Descr: sql.NullString{
					String: "",
					Valid:  false,
				},
			})
			if err != nil {
				return nil, errors.New("FAILED TO CREATE New CATEG")
			}
		}

		params := db.AddCategToBlockParams{
			BlockID: blockid,
			Name:    categ,
		}

		err = q.AddCategToBlock(ctx, params)
		if err != nil {
			log.Print("Failed to add categ to block:", err)
			return nil, errors.New("FAILED TO JOIN CATEG")
		}
	}

	publishEvent(Msg{
		Id:        blockid.String(),
		LangCode:  "",
		ChangeLog: "Added Categs",
	}, "BlockUpdated")

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        "",
		Message:   "Categ added to block",
	}, nil
}

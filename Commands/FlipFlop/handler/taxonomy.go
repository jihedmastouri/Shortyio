package handler

import (
	"context"
	"database/sql"
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
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to connect to database",
		}, nil
	}

	defer conn.Close()
	q := db.New(conn)

	params := db.CreateTagParams{
		Name:  "",
		Descr: sql.NullString{},
	}

	id, err := q.CreateTag(ctx, params)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to create tag",
		}, nil
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
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to connect to database",
		}, nil
	}

	defer conn.Close()
	q := db.New(conn)

	params := db.CreateCategParams{
		Name:  "",
		Descr: sql.NullString{},
	}

	id, err := q.CreateCateg(ctx, params)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to create category",
		}, nil
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
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to connect to database",
		}, nil
	}

	defer conn.Close()
	q := db.New(conn)

	id, err := q.DeleteTag(ctx, rq.Name)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to delete tag",
		}, nil
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
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to connect to database",
		}, nil
	}

	defer conn.Close()
	q := db.New(conn)

	id, err := q.DeleteCateg(ctx, rq.Name)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to delete category",
		}, nil
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        string(id),
		Message:   "Category deleted",
	}, nil
}

func JoinBlockTag(ctx context.Context, rq *pb.JoinTaxonomy) (*pb.ActionResponse, error) {
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

	params := db.AddTagToBlockParams{
		BlockID: blockid,
		TagID:   rq.TaxonomyId,
	}

	err = q.AddTagToBlock(ctx, params)

	if err != nil {
		log.Print("Failed to add tag to block:", err)
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   err.Error(),
		}, err
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        "",
		Message:   "Tag added to block",
	}, nil
}

func JoinBlockCategory(ctx context.Context, rq *pb.JoinTaxonomy) (*pb.ActionResponse, error) {
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

	params := db.AddCategToBlockParams{
		BlockID: blockid,
		CategID: rq.TaxonomyId,
	}

	err = q.AddCategToBlock(ctx, params)

	if err != nil {
		log.Print("Failed to add categ to block:", err)
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   err.Error(),
		}, err
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        "",
		Message:   "Categ added to block",
	}, nil
}

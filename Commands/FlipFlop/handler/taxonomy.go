package handler

import (
	"context"
	"database/sql"
	"log"

	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/flipFlop/proto"
)

// implement crud for tags and categories
// rpc CreateTag(CreateTaxonomy) returns (ActionResponse) {}
// rpc DeleteTag(DeleteTaxonomy) returns (ActionResponse) {}
// rpc JoinBlockTag(JoinTaxonomy) returns (ActionResponse) {}
//
// rpc CreateCategory(CreateTaxonomy) returns (ActionResponse) {}
// rpc DeleteCategory(DeleteTaxonomy) returns (ActionResponse) {}
// rpc JoinBlockCategory(JoinTaxonomy) returns (ActionResponse) {}

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

	id, err := q.CreateTag(ctx, params);
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

	id, err := q.CreateCateg(ctx, params);
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

	id, err := q.DeleteTag(ctx, rq.Name);
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

	id, err := q.DeleteCateg(ctx, rq.Name);
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

// update tag
func (s *CommandService) UpdateTag(ctx context.Context, in *pb.JoinTaxonomy) (*pb.ActionResponse, error) {
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

	params := db.UpdateTagParams{
		Name:  "",
		Descr: sql.NullString{},
	}

	id, err := q.UpdateTag(ctx, params);
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to update tag",
		}, nil
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        string(id),
		Message:   "Tag updated",
	}, nil
}

// update category
func (s *CommandService) UpdateCategory(ctx context.Context, rq *pb.UpdateTaxonomy) (*pb.ActionResponse, error) {
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

	params := db.UpdateCategParams{
		Name:  "",
		Descr: sql.NullString{},
	}

	id, err := q.UpdateCateg(ctx, params);
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return &pb.ActionResponse{
			IsSuceess: false,
			Id:        "",
			Message:   "Failed to update category",
		}, nil
	}

	return &pb.ActionResponse{
		IsSuceess: true,
		Id:        id,
		Message:   "Category updated",
	}, nil
}

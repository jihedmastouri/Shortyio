package handler

import (
	"context"
	"testing"

	pb "github.com/shorty-io/go-shorty/queries/proto"
)

var q = &Queries{}

func TestGetBlock(t *testing.T) {
	ctx := context.Background()

	rq := &pb.BlockRequest{
		Id:   "fc4afa73f1e7",
		Lang: "english",
	}

	res, err := q.GetBlock(ctx, rq)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res.GetMeta())
	t.Log(res.GetContent())
}

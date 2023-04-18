package handler

import (
	"context"
	"testing"

	pb "github.com/shorty-io/go-shorty/queries/proto"
)

func TestGetBlock(t *testing.T) {
	q := &Queries{}
	ctx := context.Background()

	rq := &pb.BlockRequest{
		Id: "fc4afa73f1e7",
        Lang: "english",
	}

	 res , err := q.GetBlock(ctx, rq)
     t.Log(res)
     t.Log("lol")

    if err != nil {
		t.Fatal(err)
	}
}

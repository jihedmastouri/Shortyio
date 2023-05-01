package handler
//
// import (
// 	"context"
// 	"testing"
//
// 	pb "github.com/shorty-io/go-shorty/queries/proto"
// )
//
// var (
// 	q = &Queries{}
//
// 	rq = &pb.BlockRequest{
// 		Id:   "fc4afa73f1e7",
// 		Lang: "english",
// 	}
//
// 	rqV = &pb.BlockRequest{
// 		Id:      "fc4afa73f1e7",
// 		Lang:    "english",
// 		Version: func(val int32) *int32 { return &val }(2),
// 	}
// )
//
// func TestGetBlock(t *testing.T) {
// 	ctx := context.Background()
//
// 	logResponse(q.GetBlock, t, ctx)
// 	logResponse(q.GetBlockRules, t, ctx)
// 	logResponse(q.GetBlockMeta, t, ctx)
// 	logResponse(q.GetBlockContent, t, ctx)
//
// }
//
// func TestGetBlockVersion(t *testing.T) {
// 	ctx := context.Background()
//
// 	res, err := q.GetBlock(ctx, rq)
//
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	t.Log(res)
// 	t.Log(res.GetContent())
// }
//
// func logResponse(fn func(context.Context, *pb.BlockRequest) (*any, error), t *testing.T, a context.Context) (*any, error) {
// 	res, err := fn(a, rq)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(res)
// 	return res, err
// }

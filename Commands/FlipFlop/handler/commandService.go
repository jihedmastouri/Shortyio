package handler

import (
	"context"
	"fmt"

	pb "github.com/shorty-io/go-shorty/FlipFlop/proto"
)

type commandService struct{}

func (cs *commandService) CreateCommand(ctx context.Context, rq *pb.CreateCommandRequest) (*pb.CreateCommandResponse, error) {
	resp := fmt.Sprintf("name: %v, descr: %v", rq.GetName(), rq.GetDescription())

	return &pb.CreateCommandResponse{Id: resp}, nil
}

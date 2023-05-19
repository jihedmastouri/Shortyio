package handler

import (
	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/flipFlop/proto"
)

type CommandService struct {
	pb.UnimplementedFlipFlopServer
}

// CreateBlock(CreateRequest) returns (ActionResponse) {}
// CreateBlockLang(CreateLangRequest) returns (ActionResponse) {}
// DeleteBlock(DeleteRequest) returns (ActionResponse) {}
// DeleteBlockLang(DeleteLangRequest) returns (ActionResponse) {}

func (c *CommandService) CreateBlock(*pb.CreateRequest) (*pb.ActionResponse, error) {
	return &pb.ActionResponse{
		IsSuceess: false,
		Id:        "",
		Message:   "",
	}, nil
}

func (c *CommandService) CreateBlockLang(*pb.CreateLangRequest) (*pb.ActionResponse, error) {
	return &pb.ActionResponse{
		IsSuceess: false,
		Id:        "",
		Message:   "",
	}, nil
}

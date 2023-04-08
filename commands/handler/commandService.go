package handler

import (
	pb "github.com/shorty-io/go-shorty/commands/proto"
)

type commandService struct{}

func (cs *commandService) CreateCommand(&pb.CreateCommandRequest) pb.CreateCommandResponse {

}

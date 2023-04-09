package handler

import (
	pb "github.com/shorty-io/go-shorty/FlipFlop/proto"
)

type commandService struct{}

func NewCommand()  {

}

func (cs *commandService) CreateCommand(*pb.CreateCommandRequest) (*pb.CreateCommandResponse, error) {
    return nil, nil
}

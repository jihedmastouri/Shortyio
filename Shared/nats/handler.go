package event

import (
	"time"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

func NewEvent(block *pb.Block) *BlockUpdated {
	return &BlockUpdated{
		Block: block,
		Event: Event{
			Name: "BlockUpdated",
			Time: time.Now(),
		},
	}
}

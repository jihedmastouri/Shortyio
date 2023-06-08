package event

import (
	"time"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

type Event struct {
	Name string
	time.Time
}

type BlockUpdated struct {
	*pb.Block
	Event
}

type BlockAdded struct {
	*pb.BlockMeta
	*pb.BlockRules
	Event
}

type BlockLangAdded struct {
	*pb.BlockMeta
	*pb.BlockRules
	Event
}

type UpdateBlock struct {
	*pb.BlockMeta
	*pb.BlockRules
	Event
}

type UpdateBlockContent struct {
	*pb.BlockContent
	Event
}

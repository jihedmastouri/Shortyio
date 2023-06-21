package msgq

import (
	"reflect"
	"time"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

type EventName string

const (
	BlockUpdated  EventName = "BlockUpdated"
	BlockToUpdate EventName = "BlockToUpdate"
)

type dataType interface {
	BlockToUpdateData | BlockUpdatedData
}

type Event[T dataType] struct {
	Name  EventName
	Queue string
	time.Time
	Payload T
}

type BlockUpdatedData struct {
	Id string
}

type BlockToUpdateData struct {
	Id string
	*pb.BlockContent
}

func NewEvent[T dataType](name EventName, payload T) *Event[T] {

	if name == "BlcockUpdated" {
		if reflect.TypeOf(payload) != reflect.TypeOf(BlockUpdatedData{}) {
			panic("payload must be of type BlockUpdatedData")
		}
	}

	if name == "BlockToUpdate" {
		if reflect.TypeOf(payload) != reflect.TypeOf(BlockToUpdateData{}) {
			panic("payload must be of type BlockToUpdateData")
		}
	}

	return &Event[T]{
		Name:    name,
		Queue:   "ShortyQ",
		Time:    time.Now(),
		Payload: payload,
	}
}

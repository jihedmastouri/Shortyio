package db

import (
    orm "github.com/shorty-io/go-shorty/Shared/db"
    pb "github.com/shorty-io/go-shorty/flipFlop/proto"
)

init(){


shortydb = orm.New()
}

func CreateBlock(b *pb.Block)  {
}

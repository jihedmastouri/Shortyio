package db

import (
	"database/sql"

	_ "github.com/lib/pq"

	sqlc "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/flipFlop/proto"
)

type ORM struct {
	db sqlc.DBTX
}

func New() ORM {

	ctx := context.Background()

	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		return err
	}

	queries := tutorial.New(db)
}

func CreateBlock(b *pb.Block) {
}

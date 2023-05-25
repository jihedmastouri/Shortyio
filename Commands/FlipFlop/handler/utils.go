package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/flipFlop/proto"
)

func newConn() (*sql.DB, error) {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	dbname := os.Getenv("MONGO_DBNAME")

	conn, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	))

	if err != nil {
		log.Print("Database Connection Failed", err)
		return conn, err
	}

	return conn, nil
}

// Takes a BlockRules (Name or rules) and returns a BlockRules_Rules
func getBlockRules(q *db.Queries, br *pb.BlockRules) (pb.BlockRules_Rules, string) {

	if br.GetRules() != nil {
		return *br.GetRules(), "custom"
	}

	ctx := context.Background()
	rules, err := q.GetBlockRulesByName(ctx, br.GetRuleName())
	if err != nil {
		panic(err)
	}

	return pb.BlockRules_Rules{
		Nested:            rules.Nested.Bool,
		HasLikes:          rules.HasLikes.Bool,
		HasComments:       rules.HasComments.Bool,
		CommentsNested:    rules.Nested.Bool,
		CommentsHasLikes:  rules.CommentsHasLikes.Bool,
		CommentsEditable:  rules.CommentEditable.Bool,
		CommentsMaxNested: int32(rules.CommentsMaxNest.Int16),
	}, br.GetRuleName()
}

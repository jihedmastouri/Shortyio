package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/shorty-io/go-shorty/Shared/service"

	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/flipFlop/proto"
)

var Srv *service.Service

func init() {
	Srv = service.New(service.FlipFlop)
}

func newConn() (*sql.DB, error) {
	host, err := Srv.GetKV("POSTGRES_HOST")
	port, err1 := Srv.GetKV("POSTGRES_PORT")
	user, err2 := Srv.GetKV("POSTGRES_USER")
	password, err3 := Srv.GetKV("POSTGRES_PASSWORD")
	dbname, err4 := Srv.GetKV("POSTGRES_DB")

	if err != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		log.Print("Database Connection Failed", err)
		return nil, err
	}

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

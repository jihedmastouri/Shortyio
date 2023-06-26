package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/shorty-io/go-shorty/Shared/service"

	db "github.com/shorty-io/go-shorty/Shared/db"
	pb "github.com/shorty-io/go-shorty/Shared/proto"

	_ "github.com/lib/pq"
)

var srv *service.Service
var nc *nats.Conn

func NewSrv(service *service.Service) {
	if service == nil {
		log.Fatal("Service is nil")
	}

	if srv != nil {
		log.Fatal("Service already initialized")
	}
	srv = service
}

func newConn() (*sql.DB, error) {

	params := []string{
		"POSTGRES_HOST",
		"POSTGRES_PORT",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_DB",
	}
	config := make(map[string]string)

	for _, param := range params {
		value, err := srv.GetKV(param)
		if err != nil {
			log.Fatalf(
				"Failed to retrieve %s from Consul key-value store: %s",
				param,
				err,
			)
			return nil, err
		}
		config[param] = value
	}

	// Access the parameters from the 'config' map
	host := config["POSTGRES_HOST"]
	port := config["POSTGRES_PORT"]
	user := config["POSTGRES_USER"]
	password := config["POSTGRES_PASSWORD"]
	dbname := config["POSTGRES_DB"]

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
func getBlockRules(q *db.Queries, br *pb.BlockRulesRq) pb.BlockRules {

	if br.GetRules() != nil {
		return *br.GetRules()
	}

	ctx := context.Background()
	rules, err := q.GetBlockRulesByName(ctx, br.GetRuleName())
	if err != nil {
		panic(err)
	}

	return pb.BlockRules{
		RuleName:          br.GetRuleName(),
		Nested:            rules.Nested.Bool,
		HasLikes:          rules.HasLikes.Bool,
		HasComments:       rules.HasComments.Bool,
		CommentsNested:    rules.Nested.Bool,
		CommentsHasLikes:  rules.CommentsHasLikes.Bool,
		CommentsEditable:  rules.CommentEditable.Bool,
		CommentsMaxNested: int32(rules.CommentsMaxNest.Int16),
	}
}

// func publishEvent(ctx context.Context, event service.) error {
// 	_, err := srv.Publish(ctx, event)
// 	if err != nil {
// 		log.Printf("Failed to publish event: %s", err)
// 		return err
// 	}
// 	return nil
// }

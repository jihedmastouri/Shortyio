package main

import (
	"fmt"
	"os"

	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func AggregateDB(msg Msg) (*[]byte, error) {
	query, err := os.ReadFile("./temp.sql")
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(msg.Id)
	if err != nil {
		return nil, err
	}

	data, err := executeJSONQuery(id, msg.LangCode, msg.ChangeLog, string(query))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func executeJSONQuery(id uuid.UUID, lang, changelog, query string) (*[]byte, error) {
	db, err := newConn()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var json *[]byte
	err = db.QueryRow(query, id, lang, changelog).Scan(json)
	if err != nil {
		return nil, err
	}

	return json, nil
}

func newConn() (*sql.DB, error) {
	host := os.Getenv("PG_HOST")
	port := "5432"
	user := "postgres"
	password := "root"
	dbname := "shortyio"

	conn, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	))

	if err != nil {
		return conn, err
	}

	return conn, nil
}

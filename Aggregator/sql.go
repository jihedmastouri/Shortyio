package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type foo []map[string]any

func aggregateDB(id uuid.UUID, lang string) (*foo, error) {
	query, err := os.ReadFile("./temp.sql")
	if err != nil {
		return nil, err
	}

	data, err := executeJSONQuery(id, lang, string(query))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func executeJSONQuery(id uuid.UUID, lang, query string) (*foo, error) {
	db, err := newConn()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	row := db.QueryRow(query, id, lang)

	var data string
	row.Scan(&data)
	//Scan(&json)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var res foo
	if err = json.Unmarshal([]byte(data), &res); err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

type Language struct {
	Code string `db:"lang_code"`
}

func getAllLanguages(id uuid.UUID) ([]string, error) {
	db, err := newConn()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var langs []string
	query := `SELECT lang_code FROM block_langs WHERE block_id = $1`

	rows, err := db.Query(query, id)
	if err != nil {
		return langs, err
	}
	defer rows.Close()

	for rows.Next() {
		var lang Language
		err := rows.Scan(&lang.Code)
		if err != nil {
			return langs, err
		}
		langs = append(langs, lang.Code)
	}

	return langs, nil
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

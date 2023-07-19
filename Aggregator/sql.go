package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type foo []byte

type Language struct {
	Code string `db:"lang_code"`
}

func aggregateDB(id uuid.UUID, lang string) (foo, error) {
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

func executeJSONQuery(id uuid.UUID, lang, query string) (foo, error) {
	var json []byte
	if err := db.QueryRow(query, id, lang).Scan(&json); err != nil {
		log.Println(err)
		return nil, err
	}

	return json, nil
}

func getAllLanguages(id uuid.UUID) ([]string, error) {

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

func connectSQL(config map[string]string) *sql.DB {
	host := config["POSTGRES_HOST"]
	port := config["POSTGRES_PORT"]
	user := config["POSTGRES_USER"]
	password := config["POSTGRES_PASSWORD"]
	dbname := config["POSTGRES_DB"]

	tempConn, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	))

	if err != nil {
		log.Fatal("Database Connection Failed", err)
	}

	return tempConn
}

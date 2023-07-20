package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Language struct {
	Code string `db:"lang_code"`
}

func aggregateDB(id uuid.UUID, lang string) ([]byte, error) {
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

func executeJSONQuery(id uuid.UUID, lang, query string) ([]byte, error) {
	db, err := connectSQL()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var json []byte
	if err := db.QueryRow(query, id, lang).Scan(&json); err != nil {
		return nil, fmt.Errorf("Error querying database: %v", err)
	}

	return json, nil
}

func getAllLanguages(id uuid.UUID) ([]string, error) {
	db, err := connectSQL()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT lang_code FROM block_langs WHERE block_id = $1`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("Error querying database: %v", err)
	}

	defer rows.Close()

	var langs []string
	for rows.Next() {
		var lang Language
		err := rows.Scan(&lang.Code)
		if err != nil {
			log.Println("Error scanning row: ", err)
			continue
		}
		langs = append(langs, lang.Code)
	}

	return langs, nil
}

func connectSQL() (*sql.DB, error) {
	host := globalConfig["POSTGRES_HOST"]
	port := globalConfig["POSTGRES_PORT"]
	user := globalConfig["POSTGRES_USER"]
	password := globalConfig["POSTGRES_PASSWORD"]
	dbname := globalConfig["POSTGRES_DB"]

	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	))
	if err != nil || db == nil {
		return nil, fmt.Errorf("Failed to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping the database: %v", err)
	}

	return db, nil
}

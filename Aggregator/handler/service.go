package handler

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

func AggregateDB() (string, error) {
	var conn *sql.DB
	conn, err := newConn()
	if err != nil {
		return "", err
	}

	defer conn.Close()

	query, err := os.ReadFile("../temp.sql")
	if err != nil {
		log.Print(err)
		return "", err
	}

	data, err := executeJSONQuery(conn, string(query))
	if err != nil {
		log.Print(err)
		return "", err
	}

	return data, nil
}


func executeJSONQuery(db *sql.DB, query string) ([]Data, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var people []Person

	for rows.Next() {
		var person Data
		err := rows.Scan(&person.ID, &person.Name, &person.Age)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		people = append(people, person)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during iteration: %v", err)
	}

	return people, nil
}

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

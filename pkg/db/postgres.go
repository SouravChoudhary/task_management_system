package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitPostgres() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=tasks sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}
	return db
}

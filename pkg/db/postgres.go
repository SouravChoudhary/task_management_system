package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func InitPostgres() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=tasks sslmode=disable"

	var db *sql.DB
	var err error
	// maxAttempts should be confgured from config or env var
	maxAttempts := 5
	for i := 1; i <= maxAttempts; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
		}

		if err == nil {
			log.Println("Connected to PostgreSQL")
			return db, nil
		}

		log.Printf("Failed to connect to Postgres (attempt %d/%d): %v", i, maxAttempts, err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("Postgres connection failed after %d attempts: %v", maxAttempts, err)
}

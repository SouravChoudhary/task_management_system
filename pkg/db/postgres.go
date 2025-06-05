package db

import (
	"database/sql"
	"fmt"
	"log"
	"task_management_system/config"
	"time"

	_ "github.com/lib/pq"
)

func InitPostgres(dbCnf *config.DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbCnf.Host, dbCnf.Port, dbCnf.User, dbCnf.Password, dbCnf.Name,
	)
	var db *sql.DB
	var err error

	maxAttempts := dbCnf.MaxAttempt
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

	return nil, fmt.Errorf("postgres connection failed after %d attempts: %v", maxAttempts, err)
}

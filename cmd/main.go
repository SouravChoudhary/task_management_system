package main

import (
	"database/sql"
	"log"
	"os"
	"task_management_system/config"
	"task_management_system/internal/handler"
	"task_management_system/internal/repository"
	"task_management_system/internal/service"
	"task_management_system/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config.yaml: %v", err)
	}

	postgresConn, err := db.InitPostgres(&cfg.DB)
	if err != nil {
		log.Fatalf("Unable to start service due to config error: %v", err)
	}

	err = RunMigrations(postgresConn, "./pkg/db/schema.sql")
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	taskRepo := repository.NewTaskRepository(postgresConn)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()

	taskHandler.RegisterRoutes(router)

	router.Run(":8080")
}

func RunMigrations(db *sql.DB, path string) error {
	sqlBytes, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		log.Fatalf("Failed to execute schema.sql: %v", err)
	}
	return nil
}

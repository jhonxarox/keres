package database

import (
	"database/sql"
	"log"

	"keres/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func InitDB(cfg *config.Config) *sql.DB {
	// Connect to PostgreSQL
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("PostgreSQL connected successfully!")
	return db
}

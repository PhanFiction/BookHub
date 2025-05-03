package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB // Global DB connection

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", dbUser, dbPassword, dbName)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// Optional: ping to ensure connection is valid
	if err := DB.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
	}
}

package services

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewDbConnection() *sql.DB {
	if err := godotenv.Load((".env")); err != nil {
		log.Fatalf("Failed to load .env, %v", err)
	}

	dbUrl := os.Getenv("DB_URL")
	// PostgreSQL接続設定
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil
	}
	return db
}

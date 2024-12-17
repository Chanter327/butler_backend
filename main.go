package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// PostgreSQL接続設定
	connStr := "postgres://butler:p@55vv0rd@postgres:5432/butler_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	// ヘルスチェック用
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// データベース接続確認用
	r.GET("/db-check", func(c *gin.Context) {
		err := db.Ping()
		if err != nil {
			log.Printf("Database connection error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Database connection failed",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "Database connected successfully",
		})
	})

	r.Run(":8080")
}

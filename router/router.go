package router

import (
	"log"
	"net/http"
	"os"
	"time"

	services "github.com/Chanter327/butler_backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Router() *gin.Engine {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env, %v", err)
	}
	frontUrl := os.Getenv("FRONTEND_URL")

	r := gin.Default()
	r.Use(gin.Logger())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontUrl}, // フロントエンドのURLを指定
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/api/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome to Butler",
		})
	})
	r.GET("/db-check", func(c *gin.Context) {
		err := services.NewDbConnection().Ping()
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

	return r
}
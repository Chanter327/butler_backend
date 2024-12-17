package main

import (
	"log"

	router "github.com/Chanter327/butler_backend/router"
	_ "github.com/lib/pq"
)

func main() {
	r := router.Router()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	r.Run(":8080")
}

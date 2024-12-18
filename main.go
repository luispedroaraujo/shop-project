package main

import (
	"log"
	"shop-api/config"
	"shop-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()

	r := gin.Default()

	routes.RegisterProductRoutes(r, db)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

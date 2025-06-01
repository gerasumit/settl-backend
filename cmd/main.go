package main

import (
	"fmt"
	"log"
	"os"

	"settl-backend/config"
	"settl-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load config
	config.LoadConfig()

	// Initialize routes
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

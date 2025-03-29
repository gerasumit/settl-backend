package main

import (
    "fmt"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "settl-backend/routes"
)

func main() {
    r := gin.Default()

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

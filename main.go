// main.go
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/fhrrrzy/gallery-api/middlewares"
    "github.com/fhrrrzy/gallery-api/routes"
    "github.com/fhrrrzy/gallery-api/models"
)

func main() {
    // Initialize Gin
    router := gin.Default()

    // Database setup
    db, err := gorm.Open(sqlite.Open("gallery-api.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database")
    }
    // AutoMigrate will create the necessary tables based on the provided models
    db.AutoMigrate(&models.User{}, &models.Photo{})

    // Setup routes
    routes.SetupRoutes(router)

    // Start the server
    port := 8080
    fmt.Printf("Server running on :%d\n", port)
    router.Run(fmt.Sprintf(":%d", port))
}

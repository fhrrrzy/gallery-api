// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/fhrrrzy/gallery-api/database"
	// "github.com/fhrrrzy/gallery-api/repositories"
	"github.com/fhrrrzy/gallery-api/routes"
)

func main() {
	// Initialize Gin
	router := gin.Default()

	// Database setup
	database.InitDB()
	database.MigrateDB()

	// Initialize repositories
	// repositories.Init(database.DB)

	// Setup routes
	routes.SetupRoutes(router)

	// Start the server
	port := 8080
	fmt.Printf("Server running on :%d\n", port)
	router.Run(fmt.Sprintf(":%d", port))
}

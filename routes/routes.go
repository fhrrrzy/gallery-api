// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/fhrrrzy/gallery-api/controllers"
	"github.com/fhrrrzy/gallery-api/middlewares"
)

// SetupRoutes configures API routes
func SetupRoutes(router *gin.Engine) {
	// Public routes (no authentication required)
	publicGroup := router.Group("/api")
	{
		publicGroup.POST("/users/register", controllers.RegisterUser)
		publicGroup.POST("/users/login", controllers.LoginUser)
	}

	// Authenticated routes (require authentication)
	authGroup := router.Group("/api").Use(middlewares.AuthMiddleware())
	{
		// User Endpoints
		authGroup.PUT("/users/:userId", controllers.UpdateUser)
		authGroup.DELETE("/users/:userId", controllers.DeleteUser)

		// Photos Endpoints
		authGroup.POST("/photos", controllers.CreatePhoto)
		authGroup.GET("/photos", controllers.GetPhotos)
		authGroup.PUT("/photos/:photoId", controllers.UpdatePhoto)
		authGroup.DELETE("/photos/:photoId", controllers.DeletePhoto)
	}
}

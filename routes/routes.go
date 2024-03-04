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
	publicGroup := router.Group("/")
	{
		publicGroup.POST("/users/register", controllers.RegisterUser)
		publicGroup.POST("/users/login", controllers.LoginUser)
	}

	// Authenticated routes for Users
	userAuthGroup := router.Group("/users").Use(middlewares.AuthMiddleware())
	{
		userAuthGroup.PUT("/:userId", controllers.UpdateUser)
		userAuthGroup.DELETE("/:userId", controllers.DeleteUser)
	}

	// Authenticated routes for Photos
	photoAuthGroup := router.Group("/photos").Use(middlewares.AuthMiddleware())
	{
		photoAuthGroup.POST("/", controllers.CreatePhoto)
		photoAuthGroup.GET("/", controllers.GetPhotos)
		photoAuthGroup.PUT("/:photoId", controllers.UpdatePhoto)
		photoAuthGroup.DELETE("/:photoId", controllers.DeletePhoto)
	}
}
	
// routes/routes.go
package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/fhrrrzy/gallery-api/controllers"
    "github.com/fhrrrzy/gallery-api/middlewares"
)

// SetupRoutes configures API routes
func SetupRoutes(router *gin.Engine) {
    // Apply AuthMiddleware to routes that require authorization
    authGroup := router.Group("/api").Use(middlewares.AuthMiddleware())

    // User Endpoints
    router.POST("/users/register", controllers.RegisterUser)
    router.POST("/users/login", controllers.LoginUser)
    authGroup.PUT("/users/:userId", controllers.UpdateUser)
    authGroup.DELETE("/users/:userId", controllers.DeleteUser)

    // Photos Endpoints
    authGroup.POST("/photos", controllers.CreatePhoto)
    router.GET("/photos", controllers.GetPhotos)
    authGroup.PUT("/photos/:photoId", controllers.UpdatePhoto)
    authGroup.DELETE("/photos/:photoId", controllers.DeletePhoto)
}

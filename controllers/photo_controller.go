// controllers/photo_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/fhrrrzy/gallery-api/app"
	"github.com/fhrrrzy/gallery-api/helpers"
	"github.com/fhrrrzy/gallery-api/database"
	"github.com/fhrrrzy/gallery-api/models"
)

// CreatePhoto handles creating a new photo
func CreatePhoto(c *gin.Context) {
	var createPhotoRequest app.CreatePhotoRequest
	if err := c.ShouldBindJSON(&createPhotoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract user ID from the JWT token in the Authorization header
	userIDFromToken, err := helpers.ExtractUserIDFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Convert userIDFromToken to uint
	userID := uint(userIDFromToken)

	// Create a new photo with the associated user ID
	newPhoto := models.Photo{
		Title:    createPhotoRequest.Title,
		Caption:  createPhotoRequest.Caption,
		PhotoURL: createPhotoRequest.PhotoURL,
		UserID:   userID,
	}

	// Save the new photo to the database
	if err := database.DB.Create(&newPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	// Return the created photo data in the response
	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":        newPhoto.ID,
			"title":     newPhoto.Title,
			"caption":   newPhoto.Caption,
			"photoUrl":  newPhoto.PhotoURL,
			"userId":    newPhoto.UserID,
			"createdAt": newPhoto.CreatedAt,
			"updatedAt": newPhoto.UpdatedAt,
		},
		"message": "Photo created successfully",
	})
}

// GetPhotos handles getting all photos
func GetPhotos(c *gin.Context) {
	// Perform get photos logic (you can implement this later)

	// Return a sample response for now
	c.JSON(http.StatusOK, gin.H{"message": "Get all photos"})
}

// UpdatePhoto handles updating photo information
func UpdatePhoto(c *gin.Context) {
	// Extract photo ID from path parameter
	photoID := c.Param("photoId")

	// Parse request data
	// var updatePhotoRequest UpdatePhotoRequest
	// if err := c.ShouldBindJSON(&updatePhotoRequest); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// Perform update photo logic (you can implement this later)

	// Return a sample response for now
	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully", "photoId": photoID})
}

// DeletePhoto handles deleting a photo
func DeletePhoto(c *gin.Context) {
	// Extract photo ID from path parameter
	photoID := c.Param("photoId")

	// Perform delete photo logic (you can implement this later)

	// Return a sample response for now
	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully", "photoId": photoID})
}

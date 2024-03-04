// controllers/photo_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/fhrrrzy/gallery-api/app"
	"github.com/fhrrrzy/gallery-api/helpers"
	"github.com/fhrrrzy/gallery-api/database"
	"github.com/fhrrrzy/gallery-api/models"
	"strconv"
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
	// Retrieve all photos from the database with associated user details
	var photos []models.Photo
	if err := database.DB.Preload("User").Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get photos"})
		return
	}

	// Exclude password from user details in each photo
	for i := range photos {
		photos[i].User.Password = ""
	}

	// Return the list of photos in the response
	c.JSON(http.StatusOK, gin.H{"data": photos, "message": "Photos retrieved successfully"})
}



// UpdatePhoto handles updating photo information
func UpdatePhoto(c *gin.Context) {
	// Extract photo ID from path parameter
	photoIDStr := c.Param("photoId")
	photoID, err := strconv.ParseUint(photoIDStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	// Extract the user ID from the JWT token in the Authorization header
	userIDFromToken, err := helpers.ExtractUserIDFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Check if the user is authorized to perform the update
	var existingPhoto models.Photo
	err = database.DB.First(&existingPhoto, photoID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if uint64(existingPhoto.UserID) != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}
	

	// Parse request data
	var updateRequest app.UpdatePhotoRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update photo information
	existingPhoto.Title = updateRequest.Title
	existingPhoto.Caption = updateRequest.Caption
	existingPhoto.PhotoURL = updateRequest.PhotoURL

	// Save the updated photo to the database
	if err := database.DB.Save(&existingPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	// Return a response
	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully", "photoId": photoID})
}

// DeletePhoto handles deleting a photo
func DeletePhoto(c *gin.Context) {
	// Extract photo ID from path parameter
	photoID := c.Param("photoId")

	// Parse user ID from the JWT token
	userIDFromToken, err := helpers.ExtractUserIDFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Check if the user is the owner of the photo
	var existingPhoto models.Photo
	if err := database.DB.Where("id = ?", photoID).First(&existingPhoto).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}
	if existingPhoto.UserID != uint(userIDFromToken) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	// Perform delete photo logic
	if err := database.DB.Delete(&existingPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully", "photoId": photoID})
}

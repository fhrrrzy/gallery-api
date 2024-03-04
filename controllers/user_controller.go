// controllers/user_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/fhrrrzy/gallery-api/app"
	"github.com/fhrrrzy/gallery-api/helpers"
	"github.com/fhrrrzy/gallery-api/repositories"
	"github.com/fhrrrzy/gallery-api/database"
	"strconv"
)

// UserRepository is a repository for user-related operations
var UserRepository = repositories.UserRepository{DB: database.DB}


// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	// Parse request data
	var registrationRequest app.CreateUserRequest
	if err := c.ShouldBindJSON(&registrationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate registration data
	if err := app.ValidateCreateUserRequest(registrationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// // Check if the email already exists
	// if UserRepository.EmailExists(registrationRequest.Email) {
	// 	c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
	// 	return
	// }

	// Create a new user
	newUser, err := UserRepository.CreateUserFromRequest(&registrationRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Return the registration data in the response
	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"username": newUser.Username,
			"email":    newUser.Email,
		},
		"message": "User registered successfully",
	})
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	// Parse request data
	var loginRequest app.LoginUserRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user with the provided email exists
	user, err := UserRepository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// For demonstration purposes, let's generate a sample JWT token
	token, err := helpers.GenerateJWTToken(user.Email, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	// Return the JWT token in the response
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// UpdateUser handles updating user information
func UpdateUser(c *gin.Context) {
	// Parse request data and extract user ID from path parameter
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updateRequest app.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the existing user
	existingUser, err := UserRepository.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update user information
	existingUser.Username = updateRequest.Username
	existingUser.Email = updateRequest.Email
	existingUser.Password = updateRequest.Password

	// Save the updated user
	if err := UserRepository.UpdateUser(existingUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Return a sample response for now
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "userId": userID})
}

// DeleteUser handles deleting user
func DeleteUser(c *gin.Context) {
	// Extract user ID from path parameter
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if the user with the provided ID exists
	user, err := UserRepository.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Perform user deletion logic
	if err := UserRepository.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Return a sample response for now
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "userId": userID})
}

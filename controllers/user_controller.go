// controllers/user_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/fhrrrzy/gallery-api/app"
	"github.com/fhrrrzy/gallery-api/helpers"
	"github.com/fhrrrzy/gallery-api/database"
	"github.com/fhrrrzy/gallery-api/models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	// Parse request data
	var registrationRequest app.CreateUserRequest

	// Bind JSON data
	if err := c.ShouldBindJSON(&registrationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate registration data
	if err := app.ValidateCreateUserRequest(registrationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email already exists directly in the controller
	var existingUser models.User
	err := database.DB.Where("email = ?", registrationRequest.Email).First(&existingUser).Error
	if err == nil {
		// Email already registered
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registrationRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}

	// Create a new user with the hashed password
	newUser := models.User{
		Username: registrationRequest.Username,
		Email:    registrationRequest.Email,
		Password: string(hashedPassword),
	}

	// Save the new user to the database
	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Return the registration data in the response
	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"username": newUser.Username,
			"email":    newUser.Email,
			"password": "********",
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

	// Check if the user with the provided email exists directly in the controller
	var user models.User
	err := database.DB.Where("email = ?", loginRequest.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Password is correct, generate a JWT token
	token, err := helpers.GenerateJWTToken(user.Email, user.Username, strconv.Itoa(int(user.ID)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	// Return the JWT token in the response along with other data
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"username": user.Username,
			"email":    user.Email,
			"authorization": gin.H{
				"bearer": token,
			},
		},
		"message": "Login successful",
	})
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

	// Check if the user with the provided ID exists
	var existingUser models.User
	err = database.DB.First(&existingUser, userID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update user information
	existingUser.Username = updateRequest.Username
	existingUser.Email = updateRequest.Email
	existingUser.Password = updateRequest.Password

	// Save the updated user to the database
	if err := database.DB.Save(&existingUser).Error; err != nil {
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
	var user models.User
	err = database.DB.First(&user, userID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Perform user deletion logic
	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Return a sample response for now
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "userId": userID})
}

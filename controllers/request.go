// controllers/requests.go
package controllers

// CreateUserRequest represents the request structure for user registration
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginUserRequest represents the request structure for user login
type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateUserRequest represents the request structure for updating user information
type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=6"`
}

// CreatePhotoRequest represents the request structure for creating a new photo
type CreatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photoUrl"`
	UserID   uint   `json:"userId" binding:"required"`
}

// UpdatePhotoRequest represents the request structure for updating photo information
type UpdatePhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photoUrl"`
}

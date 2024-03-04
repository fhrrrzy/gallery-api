// app/user_requests.go
package app

import "github.com/asaskevich/govalidator"

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

// ValidateCreateUserRequest validates the CreateUserRequest using govalidator
func ValidateCreateUserRequest(req CreateUserRequest) error {
	_, err := govalidator.ValidateStruct(req)
	return err
}

// ValidateLoginUserRequest validates the LoginUserRequest using govalidator
func ValidateLoginUserRequest(req LoginUserRequest) error {
	_, err := govalidator.ValidateStruct(req)
	return err
}

// ValidateUpdateUserRequest validates the UpdateUserRequest using govalidator
func ValidateUpdateUserRequest(req UpdateUserRequest) error {
	_, err := govalidator.ValidateStruct(req)
	return err
}

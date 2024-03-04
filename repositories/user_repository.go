// repositories/user_repository.go
package repositories

import (
	"github.com/fhrrrzy/gallery-api/models"
	"gorm.io/gorm"
	"github.com/fhrrrzy/gallery-api/app"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(user *models.User) error {
	return ur.DB.Create(user).Error
}

// GetUserByEmail retrieves a user by email from the database
func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// GetUserByID retrieves a user by ID from the database
func (ur *UserRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := ur.DB.First(&user, userID).Error
	return &user, err
}

// UpdateUser updates user information in the database
func (ur *UserRepository) UpdateUser(user *models.User) error {
	return ur.DB.Save(user).Error
}

// DeleteUser deletes a user from the database
func (ur *UserRepository) DeleteUser(user *models.User) error {
	return ur.DB.Delete(user).Error
}

// EmailExists checks if an email already exists in the database
func (ur *UserRepository) EmailExists(email string) bool {
	var count int64
	ur.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// CreateUserFromRequest creates a new user in the database using registration request data
func (ur *UserRepository) CreateUserFromRequest(registrationRequest *app.CreateUserRequest) (*models.User, error) {
	newUser := models.User{
		Username: registrationRequest.Username,
		Email:    registrationRequest.Email,
		Password: registrationRequest.Password,
	}

	if err := ur.CreateUser(&newUser); err != nil {
		return nil, err
	}

	return &newUser, nil
}
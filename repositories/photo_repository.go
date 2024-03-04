// repositories/photo_repository.go
package repositories

import (
	"github.com/fhrrrzy/gallery-api/models" // Assuming your models are in a separate package
	"gorm.io/gorm"
)


type PhotoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{DB: db}
}


// CreatePhoto creates a new photo in the database
func (r *PhotoRepository) CreatePhoto(photo *models.Photo) error {
	return r.DB.Create(photo).Error
}

// GetPhotoByID retrieves a photo by its ID
func (r *PhotoRepository) GetPhotoByID(photoID uint) (*models.Photo, error) {
	var photo models.Photo
	err := r.DB.First(&photo, photoID).Error
	return &photo, err
}

// UpdatePhoto updates an existing photo in the database
func (r *PhotoRepository) UpdatePhoto(photo *models.Photo) error {
	return r.DB.Save(photo).Error
}

// DeletePhoto deletes a photo from the database
func (r *PhotoRepository) DeletePhoto(photoID uint) error {
	return r.DB.Delete(&models.Photo{}, photoID).Error
}

// GetPhotosByUserID retrieves all photos associated with a user
func (r *PhotoRepository) GetPhotosByUserID(userID uint) ([]models.Photo, error) {
	var photos []models.Photo
	err := r.DB.Where("user_id = ?", userID).Find(&photos).Error
	return photos, err
}

// Other methods as needed...

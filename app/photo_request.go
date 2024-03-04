// app/photo_requests.go
package app

import "github.com/asaskevich/govalidator"

// CreatePhotoRequest represents the request structure for creating a new photo
type CreatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photoUrl"`
}

// UpdatePhotoRequest represents the request structure for updating photo information
type UpdatePhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photoUrl"`
}

// ValidateCreatePhotoRequest validates the CreatePhotoRequest using govalidator
func ValidateCreatePhotoRequest(req CreatePhotoRequest) error {
	_, err := govalidator.ValidateStruct(req)
	return err
}

// ValidateUpdatePhotoRequest validates the UpdatePhotoRequest using govalidator
func ValidateUpdatePhotoRequest(req UpdatePhotoRequest) error {
	_, err := govalidator.ValidateStruct(req)
	return err
}

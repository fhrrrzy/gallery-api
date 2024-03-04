// repositories/repositories.go
package repositories

import "gorm.io/gorm"

var (
	UserRepository *UserRepository
	PhotoRepository *PhotoRepository
)

// Init initializes repositories with the provided database instance
func Init(db *gorm.DB) {
	UserRepository = NewUserRepository(db)
	PhotoRepository = NewPhotoRepository(db)
}

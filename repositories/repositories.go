// repositories/repositories.go
package repositories

import "gorm.io/gorm"

// Init initializes repositories with the provided database instance
func Init(db *gorm.DB) {
	UserRepository := UserRepository{DB: db}
	PhotoRepository := PhotoRepository{DB: db}

	// Assign the initialized repositories to the global variables
	UserRepository = UserRepository
	PhotoRepository = PhotoRepository
}

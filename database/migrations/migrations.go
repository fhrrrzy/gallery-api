// database/migrations/migrations.go
package migrations

import (
	"gorm.io/gorm"
	"github.com/fhrrrzy/gallery-api/models"
)

// Migrate runs auto migration on the provided database connection
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Photo{})
}

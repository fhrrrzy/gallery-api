// database/database.go
package database

import (
    "fmt"
    "github.com/fhrrrzy/gallery-api/database/migrations"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
    var err error
    dsn := "root:@tcp(127.0.0.1:3306)/gallery-db?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database")
    }

    fmt.Println("Connected to the database")

    // Uncomment the line below if you want to log SQL statements
    // DB = DB.Debug()
}

// MigrateDB migrates the database schema
func MigrateDB() {
    // AutoMigrate will create the necessary tables based on the provided models
    migrations.Migrate(DB)
    fmt.Println("Database migrated")
}

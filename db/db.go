package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"meme-coin-api/models"
)

var DB *gorm.DB

// InitDB initializes the production database
func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("meme_coins.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return DB.AutoMigrate(&models.MemeCoin{})
}

// InitTestDB initializes an in-memory SQLite database for testing
func InitTestDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return err
	}
	return DB.AutoMigrate(&models.MemeCoin{})
}

// CloseTestDB closes the test database connection
func CloseTestDB() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}

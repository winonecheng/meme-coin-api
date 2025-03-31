package db

import (
	"meme-coin-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("meme_coins.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate the database schema
	DB.AutoMigrate(&models.MemeCoin{})
}

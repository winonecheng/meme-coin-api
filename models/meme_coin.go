package models

import "time"

// MemeCoin represents the data structure for a meme coin
type MemeCoin struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"unique;not null"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	PopularityScore int       `json:"popularity_score"`
}

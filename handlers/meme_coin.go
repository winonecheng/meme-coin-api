package handlers

import (
	"meme-coin-api/db"
	"meme-coin-api/models"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateMemeCoin handles the creation of a new meme coin
func CreateMemeCoin(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	memeCoin := models.MemeCoin{
		Name:            input.Name,
		Description:     input.Description,
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}

	if result := db.DB.Create(&memeCoin); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create meme coin"})
		return
	}

	c.JSON(http.StatusCreated, memeCoin)
}

// GetMemeCoin retrieves details of a meme coin by ID
func GetMemeCoin(c *gin.Context) {
	id := c.Param("id")
	var memeCoin models.MemeCoin

	if result := db.DB.First(&memeCoin, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meme coin not found"})
		return
	}

	c.JSON(http.StatusOK, memeCoin)
}

// UpdateMemeCoin updates the description of a meme coin by ID
func UpdateMemeCoin(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var memeCoin models.MemeCoin
	if result := db.DB.First(&memeCoin, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meme coin not found"})
		return
	}

	memeCoin.Description = input.Description
	db.DB.Save(&memeCoin)

	c.JSON(http.StatusOK, memeCoin)
}

// DeleteMemeCoin removes a meme coin by ID
func DeleteMemeCoin(c *gin.Context) {
	id := c.Param("id")
	if result := db.DB.Delete(&models.MemeCoin{}, id); result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meme coin not found"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Meme coin deleted"})
}

// PokeMemeCoin increments the popularity score of a meme coin by ID
func PokeMemeCoin(c *gin.Context) {
	id := c.Param("id")
	var memeCoin models.MemeCoin

	if result := db.DB.First(&memeCoin, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meme coin not found"})
		return
	}

	memeCoin.PopularityScore++
	db.DB.Save(&memeCoin)

	c.JSON(http.StatusOK, memeCoin)
}

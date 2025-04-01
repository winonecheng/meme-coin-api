package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"meme-coin-api/db"
	"meme-coin-api/models"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Initialize an in-memory SQLite database for testing
	db.InitTestDB()

	// Register routes
	router.POST("/meme-coins", CreateMemeCoin)
	router.GET("/meme-coins/:id", GetMemeCoin)
	router.PUT("/meme-coins/:id", UpdateMemeCoin)
	router.DELETE("/meme-coins/:id", DeleteMemeCoin)
	router.POST("/meme-coins/:id/poke", PokeMemeCoin)

	return router
}

func TestCreateMemeCoin(t *testing.T) {
	router := setupTestRouter()
	defer db.CloseTestDB()

	payload := map[string]string{
		"name":        "DogeCoin",
		"description": "A fun cryptocurrency",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/meme-coins", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.MemeCoin
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "DogeCoin", response.Name)
	assert.Equal(t, "A fun cryptocurrency", response.Description)
}

func TestGetMemeCoin(t *testing.T) {
	router := setupTestRouter()
	defer db.CloseTestDB()

	// Seed the database with a test record
	memeCoin := models.MemeCoin{Name: "DogeCoin", Description: "A fun cryptocurrency"}
	db.DB.Create(&memeCoin)

	req, _ := http.NewRequest("GET", "/meme-coins/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.MemeCoin
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "DogeCoin", response.Name)
	assert.Equal(t, "A fun cryptocurrency", response.Description)
}

func TestUpdateMemeCoin(t *testing.T) {
	router := setupTestRouter()
	defer db.CloseTestDB()

	// Seed the database with a test record
	memeCoin := models.MemeCoin{Name: "DogeCoin", Description: "A fun cryptocurrency"}
	db.DB.Create(&memeCoin)

	payload := map[string]string{
		"description": "An updated description",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT", "/meme-coins/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.MemeCoin
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "An updated description", response.Description)
}

func TestDeleteMemeCoin(t *testing.T) {
	router := setupTestRouter()
	defer db.CloseTestDB()

	// Seed the database with a test record
	memeCoin := models.MemeCoin{Name: "DogeCoin", Description: "A fun cryptocurrency"}
	db.DB.Create(&memeCoin)

	req, _ := http.NewRequest("DELETE", "/meme-coins/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Verify the record is deleted
	var deletedMemeCoin models.MemeCoin
	result := db.DB.First(&deletedMemeCoin, 1)
	assert.Error(t, result.Error)
}

func TestPokeMemeCoin(t *testing.T) {
	router := setupTestRouter()
	defer db.CloseTestDB()

	// Seed the database with a test record
	memeCoin := models.MemeCoin{Name: "DogeCoin", Description: "A fun cryptocurrency", PopularityScore: 10}
	db.DB.Create(&memeCoin)

	req, _ := http.NewRequest("POST", "/meme-coins/1/poke", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.MemeCoin
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 11, response.PopularityScore)
}

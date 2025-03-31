package main

import (
	"meme-coin-api/db"
	"meme-coin-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Set up Gin router
	r := gin.Default()

	// Define API routes
	r.POST("/meme-coins", handlers.CreateMemeCoin)
	r.GET("/meme-coins/:id", handlers.GetMemeCoin)
	r.PUT("/meme-coins/:id", handlers.UpdateMemeCoin)
	r.DELETE("/meme-coins/:id", handlers.DeleteMemeCoin)
	r.POST("/meme-coins/:id/poke", handlers.PokeMemeCoin)

	// Start the server
	r.Run(":8080")
}

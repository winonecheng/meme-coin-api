package main

import (
	"github.com/gin-gonic/gin"

	"meme-coin-api/db"
	"meme-coin-api/handlers"
)

func main() {
	// Initialize the database
	if err := db.InitDB(); err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

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

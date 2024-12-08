package main

import (
	"log"
	"net/http"

	"keres/config"
	"keres/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db := database.InitDB(cfg)
	defer db.Close()

	// Set up Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Application is running!",
		})
	})

	// Start the server
	log.Println("Starting server on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

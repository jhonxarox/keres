package main

import (
	"keres/config"
	"keres/internal/customer"
	"keres/internal/database"
	"keres/internal/limit"
	"keres/internal/transaction"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations(dsn string) {
	m, err := migrate.New(
		"file://migrations", // Path to the migrations folder
		dsn,
	)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Run database migrations
	runMigrations(cfg.DSN())

	// Initialize database
	db := database.InitDB(cfg)
	defer db.Close()

	// Initialize repositories and handlers
	customerRepo := customer.NewRepository(db)
	customerHandler := customer.NewHandler(customerRepo)

	transactionRepo := transaction.NewRepository(db)
	transactionHandler := transaction.NewHandler(transactionRepo)

	limitRepo := limit.NewRepository(db)
	limitHandler := limit.NewHandler(limitRepo)

	// Set up Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Application is running!",
		})
	})

	// Customer routes
	r.GET("/customers", customerHandler.GetAllCustomers)
	r.POST("/customers", customerHandler.CreateCustomer)

	// Transaction routes
	r.POST("/transactions", transactionHandler.CreateTransaction)

	// Limit routes
	r.GET("/limits/:customer_id", limitHandler.GetLimits)
	r.POST("/limits", limitHandler.CreateLimit)

	// Start the server
	log.Println("Starting server on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

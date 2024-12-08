package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func LoadConfig() *Config {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file. Falling back to system environment variables.")
	}

	return &Config{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}

func (c *Config) DSN() string {
	// PostgreSQL connection string
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUsername, c.DBPassword, c.DBName)
}

package config

import (
	"log"
	"os"

	"github.com/ArbinBhasimaOfficial/JobTracker/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Read the single connection URI directly
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is missing")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Neon database: %v", err)
	}

	// Auto-migration runs on the cloud instance automatically
	err = database.AutoMigrate(&models.Application{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	DB = database
}

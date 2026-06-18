package main

import (
	"log"
	"os"

	"github.com/ArbinBhasimaOfficial/JobTracker/config"
	"github.com/ArbinBhasimaOfficial/JobTracker/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	config.ConnectDB()

	r := routes.SetupRouter()

	r.SetTrustedProxies([]string{"127.0.0.1"})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

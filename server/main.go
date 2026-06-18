package main

import (
	"log"
	"os"

	"github.com/ArbinBhasimaOfficial/JobTracker/config"
	"github.com/ArbinBhasimaOfficial/JobTracker/routes"
	"github.com/gin-contrib/cors"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	config.ConnectDB()

	r := routes.SetupRouter()

	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://your-vercel-app.vercel.app",
		},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

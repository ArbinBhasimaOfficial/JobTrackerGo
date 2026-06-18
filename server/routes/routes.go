package routes

import (
	"github.com/ArbinBhasimaOfficial/JobTracker/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Strict CORS settings matching a modern production local environment
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	r.Use(cors.New(config))

	api := r.Group("/applications")
	{
		api.GET("", controllers.GetApplications)
		api.GET("/:id", controllers.GetApplicationByID)
		api.POST("", controllers.CreateApplication)
		api.PATCH("/:id", controllers.UpdateApplication)
		api.DELETE("/:id", controllers.DeleteApplication)
	}

	return r
}

package routes

import (
	"github.com/ArbinBhasimaOfficial/JobTracker/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://job-tracker-go-nextjs-9xqk-4fr68v2rr.vercel.app",
			"https://job-tracker-go-nextjs-9xqk-d4ath4ndh.vercel.app",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
		},
	}))

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

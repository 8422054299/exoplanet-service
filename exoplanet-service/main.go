package main

import (
	"exoplanet-service/database"
	"exoplanet-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	r.POST("/exoplanets", handlers.CreateExoplanet)
	r.GET("/exoplanets", handlers.GetExoplanets)
	r.GET("/exoplanets/:id", handlers.GetExoplanetByID)
	r.PUT("/exoplanets/:id", handlers.UpdateExoplanet)
	r.DELETE("/exoplanets/:id", handlers.DeleteExoplanet)
	r.GET("/fuel-estimation", handlers.FuelEstimation)

	r.Run(":8080")
}

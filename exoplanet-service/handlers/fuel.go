package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FuelEstimation(c *gin.Context) {
	distance, err := strconv.Atoi(c.Query("distance"))
	if err != nil || distance < 10 || distance > 1000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid distance"})
		return
	}

	radius, err := strconv.ParseFloat(c.Query("radius"), 64)
	if err != nil || radius <= 0.1 || radius >= 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid radius"})
		return
	}

	crewCapacity, err := strconv.Atoi(c.Query("crewCapacity"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid crew capacity"})
		return
	}

	planetType := c.Query("type")
	var gravity float64

	switch planetType {
	case "GasGiant":
		gravity = 0.5 / (radius * radius)
	case "Terrestrial":
		mass, err := strconv.ParseFloat(c.Query("mass"), 64)
		if err != nil || mass <= 0.1 || mass >= 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mass"})
			return
		}
		gravity = mass / (radius * radius)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exoplanet type"})
		return
	}

	fuel := float64(distance) / (gravity * gravity) * float64(crewCapacity)
	c.JSON(http.StatusOK, gin.H{"fuel": fuel})
}

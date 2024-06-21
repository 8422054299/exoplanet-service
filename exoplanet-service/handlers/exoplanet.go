package handlers

import (
	"exoplanet-service/database"
	"exoplanet-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func CreateExoplanet(c *gin.Context) {
	var input models.Exoplanet
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetExoplanets(c *gin.Context) {
	var exoplanets []models.Exoplanet
	database.DB.Find(&exoplanets)

	c.JSON(http.StatusOK, gin.H{"data": exoplanets})
}

func GetExoplanetByID(c *gin.Context) {
	var exoplanet models.Exoplanet
	id := c.Param("id")

	if err := database.DB.Where("id = ?", id).First(&exoplanet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exoplanet not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": exoplanet})
}

func UpdateExoplanet(c *gin.Context) {
	var exoplanet models.Exoplanet
	id := c.Param("id")

	if err := database.DB.Where("id = ?", id).First(&exoplanet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exoplanet not found!"})
		return
	}

	var input models.Exoplanet
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&exoplanet).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": exoplanet})
}

func DeleteExoplanet(c *gin.Context) {
	var exoplanet models.Exoplanet
	id := c.Param("id")

	if err := database.DB.Where("id = ?", id).First(&exoplanet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exoplanet not found!"})
		return
	}

	database.DB.Delete(&exoplanet)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

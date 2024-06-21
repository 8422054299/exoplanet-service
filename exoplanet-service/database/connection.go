package database

import (
	"exoplanet-service/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(localhost:3306)/exoplanet_service?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Exoplanet{})

	DB = database
}

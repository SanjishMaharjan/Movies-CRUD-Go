package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"src/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!!!")
	}
	DB.AutoMigrate(&models.Movie{})
}

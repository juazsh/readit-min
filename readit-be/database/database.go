package database

import (
	"log"
	"readit-be/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var err error
	DB, err = gorm.Open(sqlite.Open("readit.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database! : %v", err)
	}

	DB.AutoMigrate(&models.Post{}, &models.User{})
}

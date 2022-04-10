package db

import (
	"log"

	"github.com/kinduff/tech_qa/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("./data/database.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	log.Println("Connected to database")

	database.AutoMigrate(&models.Question{})

	DB = database
}

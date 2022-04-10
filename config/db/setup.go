package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/kinduff/tech_qa/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("./data/database.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	log.Println("Connected to database")

	DB = database
}

func DropDB(db *gorm.DB) {
	log.Println("Dropping database...")
	db.Migrator().DropTable(&models.Question{})
}

func CreateDB(db *gorm.DB) {
	log.Println("Creating database...")
	db.AutoMigrate(&models.Question{})
}

func SetupDB(db *gorm.DB) {
	DropDB(db)
	CreateDB(db)
}

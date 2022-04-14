package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/kinduff/techq/internal/models"
)

// DB is the database connection.
var DB *gorm.DB

// ConnectDatabase method connects to the database.
func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("./data/database.db"), &gorm.Config{})

	if err != nil {
		log.WithField("event", "database").Fatal("Failed to connect to database!")
	}

	log.WithField("event", "database").Info("Connected to database")

	DB = database
}

// DropDB method drops the database.
func DropDB(db *gorm.DB) {
	log.WithField("event", "database").Info("Dropping database")
	err := db.Migrator().DropTable(&models.Question{})
	if err != nil {
		log.WithField("event", "database").Fatal("Failed to drop database")
	}
}

// CreateDB method creates the database.
func CreateDB(db *gorm.DB) {
	log.WithField("event", "database").Info("Creating database")
	err := db.AutoMigrate(&models.Question{})
	if err != nil {
		log.WithField("event", "database").Fatal("Failed to create database")
	}
}

// SetupDB drops and creates the database.
func SetupDB(db *gorm.DB) {
	DropDB(db)
	CreateDB(db)
}

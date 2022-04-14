package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/kinduff/techq/db"
	log "github.com/sirupsen/logrus"
)

// Conf is the global configuration.
var Conf *config

type config struct {
	Port string
}

// LoadConfig loads the configuration.
func LoadConfig() {
	loadDotEnv()
	Conf = &config{
		Port: getEnv("PORT", "3000"),
	}
}

// HandleArgs handles the command line arguments.
func HandleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			db.ExecuteSeed(db.DB, args[1:]...)
			os.Exit(0)
		case "create":
			db.CreateDB(db.DB)
			os.Exit(0)
		case "drop":
			db.DropDB(db.DB)
			os.Exit(0)
		case "setup":
			db.SetupDB(db.DB)
			db.ExecuteSeed(db.DB)
			os.Exit(0)
		}
	}
}

func loadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.WithField("event", "dotenv").Info("No .env file found")
	} else {
		log.WithField("event", "dotenv").Info("Loaded .env file")
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

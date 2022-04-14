package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/kinduff/tech_qa/config"
	"github.com/kinduff/tech_qa/db"
	"github.com/kinduff/tech_qa/internal/server"
)

var (
	s *server.Server

	// Conf is the global configuration.
	Conf *config.Config
)

func main() {
	loadDotEnv()
	loadConfig()
	db.ConnectDatabase()
	handleArgs()
	initHTTPServer(Conf.Port)
	handleExitSignal()
}

func initHTTPServer(port string) {
	s = server.NewServer(port)
	go s.ListenAndServe()
}

func handleExitSignal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	s.Stop()
	log.WithField("event", "server").Fatal("HTTP server stopped")
}

func handleArgs() {
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

func loadConfig() {
	Conf = config.New()
}

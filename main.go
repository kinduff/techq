package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/kinduff/tech_qa/config"
	"github.com/kinduff/tech_qa/db"
	"github.com/kinduff/tech_qa/internal/server"
)

var (
	s *server.Server
)

func main() {
	config.LoadConfig()
	db.ConnectDatabase()
	config.HandleArgs()
	initHTTPServer(config.Conf.Port)
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

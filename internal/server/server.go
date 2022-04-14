// Package server provides an HTTP server to serve the metrics and other endpoints.
package server

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kinduff/techq/internal/handlers"
)

// Server is an HTTP server.
type Server struct {
	httpServer *http.Server
}

// NewServer method initializes a new HTTP server instance and uses the defined handlers.
func NewServer(port string) *Server {
	mux := http.NewServeMux()
	httpServer := &http.Server{Addr: ":" + port, Handler: mux}

	s := &Server{
		httpServer: httpServer,
	}

	mux.HandleFunc("/", withLogging(handlers.IndexHandler))
	mux.HandleFunc("/q/", withLogging(handlers.ShowHandler))
	mux.Handle("/static/", handlers.StaticHandler())

	return s
}

// ListenAndServe method serves HTTP requests.
func (s *Server) ListenAndServe() {
	log.WithFields(log.Fields{
		"addr":  "http://localhost" + s.httpServer.Addr,
		"event": "server",
	}).Info("Starting HTTP server")

	if err := s.httpServer.ListenAndServe(); err != nil {
		log.WithField("event", "server").Fatal(err)
	}
}

// Stop method stops the HTTP server.
func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.WithField("event", "server").Fatal(err)
	}
}

func withLogging(h http.HandlerFunc) http.HandlerFunc {
	logFn := func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()

		uri := r.RequestURI
		method := r.Method
		h.ServeHTTP(rw, r)

		duration := time.Since(start)

		log.WithFields(log.Fields{
			"uri":      uri,
			"method":   method,
			"duration": duration,
			"event":    "request",
		}).Info("Request served")
	}
	return http.HandlerFunc(logFn)
}

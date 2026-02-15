package main

import (
	"log"
	"video-server/internal/config"
	"video-server/internal/server"
)

func main() {
	log.Println("Initializing Video Server...")

	// Load configuration
	cfg := config.Load()

	// Initialize server
	srv := server.New(cfg)

	// Start server
	// This will block until the server is stopped (which is currently not gracefully handled, but acceptable for MVP)
	srv.Start()
}

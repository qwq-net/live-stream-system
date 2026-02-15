package main

import (
	"log"
	"video-server/internal/config"
	"video-server/internal/server"
)

func main() {
	log.Println("Initializing Video Server...")

	cfg := config.Load()

	srv := server.New(cfg)

	srv.Start()
}

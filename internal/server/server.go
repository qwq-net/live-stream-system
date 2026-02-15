package server

import (
	"log"
	"net/http"
	"video-server/internal/config"
	"video-server/internal/stream"

	"github.com/nareix/joy4/format/rtmp"
)

// Server represents the streaming server.
type Server struct {
	config *config.Config
	stream *stream.Handler
}

// New creating a new server instance.
func New(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
		stream: stream.NewHandler(),
	}
}

// Start starts the RTMP and HTTP servers.
// This method blocks.
func (s *Server) Start() {
	// RTMP Server
	rtmpServer := &rtmp.Server{
		Addr: ":" + s.config.RTMPPort,
	}

	rtmpServer.HandlePublish = s.stream.HandleRTMPPublish

	go func() {
		log.Printf("Starting RTMP server on port %s", s.config.RTMPPort)
		if err := rtmpServer.ListenAndServe(); err != nil {
			log.Fatalf("RTMP server error: %v", err)
		}
	}()

	// HTTP Server
	// Serve static files from web directory
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	http.HandleFunc("/live/", s.stream.HandleHTTPPlay)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("Starting HTTP server on port %s", s.config.HTTPPort)
	if err := http.ListenAndServe(":"+s.config.HTTPPort, nil); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}

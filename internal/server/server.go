package server

import (
	"log"
	"net/http"
	"video-server/internal/config"
	"video-server/internal/stream"

	"github.com/nareix/joy4/format/rtmp"
)

type Server struct {
	config *config.Config
	stream *stream.Handler
}

func New(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
		stream: stream.NewHandler(),
	}
}

func (s *Server) Start() {
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

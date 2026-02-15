package stream

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format/flv"
	"github.com/nareix/joy4/format/rtmp"
)

// Handler manages video streams.
type Handler struct {
	queues map[string]*pubsub.Queue
	mutex  sync.RWMutex
}

// NewHandler creates a new stream handler.
func NewHandler() *Handler {
	return &Handler{
		queues: make(map[string]*pubsub.Queue),
	}
}

// HandleRTMPPublish handles RTMP publish requests.
func (h *Handler) HandleRTMPPublish(conn *rtmp.Conn) {
	streams, _ := conn.Streams()
	path := conn.URL.Path

	// Normalize path key (e.g., /live/test key)
	key := strings.TrimPrefix(path, "/")

	queue := pubsub.NewQueue()
	queue.WriteHeader(streams)

	h.mutex.Lock()
	h.queues[key] = queue
	h.mutex.Unlock()

	log.Printf("Stream published: %s", key)

	avutil.CopyFile(queue, conn)

	h.mutex.Lock()
	delete(h.queues, key)
	h.mutex.Unlock()

	queue.Close()
	log.Printf("Stream closed: %s", key)
}

// HandleHTTPPlay handles HTTP-FLV playback requests.
func (h *Handler) HandleHTTPPlay(w http.ResponseWriter, r *http.Request) {
	// Path example: /live/test.flv
	path := r.URL.Path

	// key should match the publish key.
	// If publish is /live/test, key is live/test
	// basic mapping logic: remove leading slash and .flv extension
	key := strings.TrimPrefix(path, "/")
	key = strings.TrimSuffix(key, ".flv")

	h.mutex.RLock()
	queue, exists := h.queues[key]
	h.mutex.RUnlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "video/x-flv")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cursor := queue.Latest()
	muxer := flv.NewMuxer(w)

	avutil.CopyFile(muxer, cursor)
}

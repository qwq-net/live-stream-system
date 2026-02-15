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

type Handler struct {
	queues map[string]*pubsub.Queue
	mutex  sync.RWMutex
}

func NewHandler() *Handler {
	return &Handler{
		queues: make(map[string]*pubsub.Queue),
	}
}

func (h *Handler) HandleRTMPPublish(conn *rtmp.Conn) {
	streams, _ := conn.Streams()
	path := conn.URL.Path

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

func (h *Handler) HandleHTTPPlay(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

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

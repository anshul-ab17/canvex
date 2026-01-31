package http

import (
	"net/http"

	"server/core/ws"
)

func NewRouter(hub *ws.Hub) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/rooms", Rooms)
	mux.HandleFunc("/rooms/", GetRoom)

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWS(hub, w, r)
	})

	var handler http.Handler = mux
	handler = Logger(handler)

	return handler
}

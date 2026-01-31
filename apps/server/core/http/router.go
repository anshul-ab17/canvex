package http

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/rooms", Rooms)
	mux.HandleFunc("/rooms/", GetRoom)


	var handler http.Handler = mux
	handler = Logger(handler)

	return handler
}

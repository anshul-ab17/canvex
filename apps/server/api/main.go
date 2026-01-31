package main

import (
	"log"
	"net/http"

	httpLayer "server/core/http"
	"server/core/ws"
)
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	hub := ws.NewHub()
	go hub.Run()

	router := httpLayer.NewRouter(hub)

	log.Println("Canvex running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"log"
	"net/http"

	httpLayer "server/core/http"
	"server/core/ws"
)

func main() {
	hub := ws.NewHub()
	go hub.Run()

	router := httpLayer.NewRouter(hub)

	log.Println("🚀 Canvex API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

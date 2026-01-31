package main

import (
	"log"
	"net/http"

	"server/core/ws"
)

func main() {
	hub := ws.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWS(hub, w, r)
	})

	log.Println("WS server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

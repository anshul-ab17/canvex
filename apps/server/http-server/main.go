package main

import (
	"log"
	"net/http"

	httpLayer "server/core/http"
)
func main() {
	router := httpLayer.NewRouter()

	log.Println("Canvex running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

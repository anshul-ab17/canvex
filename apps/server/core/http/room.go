package http

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Rooms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	roomID := "room123"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"id": roomID,
		"wsUrl": "ws://localhost:8081/ws?roomId=" + roomID,
	})
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/rooms/")
	if id == "" {
		http.Error(w, "room id required", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"id": id,
	})
}

package ws

import "net/http"

func ServeWS(_ *Hub, w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

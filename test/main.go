package main

import (
	"encoding/json"
	"net/http"

	"github.com/NameLessCorporation/live-chat-lib/hub"
	"github.com/NameLessCorporation/live-chat-lib/models"
	websoket "github.com/NameLessCorporation/live-chat-lib/websocket"
)

func main() {
	handler()
	http.ListenAndServe(":8080", nil)
}

func handler() {
	h := hub.NewHub()
	go h.Run()
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		ws := websoket.NewWebSocket(h)
		var clientInfo hub.ClientInfo
		body := NewRequestReader(r)
		json.Unmarshal(body, &clientInfo)
		err := ws.RunWebSocket(&models.WebSocketConfig{
			Response: w,
			Request:  *r,
			Token:    "12345678",
		}, &clientInfo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		b, err := json.Marshal(err)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		NewResponseWriter(b, w)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
}

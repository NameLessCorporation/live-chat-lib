package main

import (
	"log"
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
		clientInfo := hub.ClientInfo{
			Name:  "Igor",
			Email: "dfdfd",
			Token: "12345678",
		}
		err := ws.RunWebSocket(&models.WebSocketConfig{
			Response: w,
			Request:  *r,
			Token:    "12345678",
		}, &clientInfo)
		if err != nil {
			log.Println(err)
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
}

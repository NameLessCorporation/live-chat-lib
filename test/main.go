package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NameLessCorporation/live-chat-lib/hub"
	models "github.com/NameLessCorporation/live-chat-lib/models"
)

func main() {
	handler()
	http.ListenAndServe(":8080", nil)
}

func handler() {
	h := hub.NewHub()
	go h.Run()
	var clientInfo hub.ClientInfo
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		ws := websoket.NewWebSocket(h)
		ws.RunWebSocket(&models.WebSocketConfig{
			Response: w,
			Request:  *r,
			Token:    "12345678",
		}, &clientInfo)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Request error: ", err)
		}
		json.Unmarshal(body, &clientInfo)
	})
}

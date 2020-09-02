package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NameLessCorporation/live-chat-lib/hub"
	"github.com/NameLessCorporation/live-chat-lib/models"
	websocket "github.com/NameLessCorporation/live-chat-lib/websocket"
)

func main() {
	handler()
	http.ListenAndServe(":8080", nil)
}

func handler() {
	rooms := models.NewRooms()
	// var ClientsQueue []*hub.ClientInfo
	var clientInfo hub.ClientInfo
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		ws := websocket.NewWebSocket(w, r)
		ws.ConnectionWebSocket(&clientInfo, rooms)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			Name  string `json:"name"`
			Token string `json:"token"`
		}
		var req Request
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Request error: ", err)
		}
		json.Unmarshal(body, &req)
		room := &models.Room{
			Name:    req.Name,
			Token:   req.Token,
			Clients: nil,
			Hub:     nil,
		}
		rooms.Create(room)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			Token string `json:"token"`
		}
		var req Request
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Request error: ", err)
		}
		json.Unmarshal(body, &req)
		room := &models.Room{
			Token: req.Token,
		}
		rooms.Delete(room)
	})

	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Request error: ", err)
		}
		json.Unmarshal(body, &clientInfo)
	})
}

package websoket

import (
	"log"
	"net/http"

	"github.com/NameLessCorporation/live-chat-lib/hub"
	"github.com/NameLessCorporation/live-chat-lib/models"
	"github.com/gorilla/websocket"
)

// WebSocket ...
type WebSocket struct {
	Upgrader *websocket.Upgrader
	Response http.ResponseWriter
	Request  *http.Request
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// NewWebSocket ...
func NewWebSocket(w http.ResponseWriter, r *http.Request) *WebSocket {
	return &WebSocket{
		Upgrader: &upgrader,
		Response: w,
		Request:  r,
	}
}

// ConnectionWebSocket ...
func (ws *WebSocket) ConnectionWebSocket(clientInfo *hub.ClientInfo, rooms *models.Rooms) error {
	for _, room := range rooms.Rooms {
		if room.Token == clientInfo.Token {
			conn, err := ws.Upgrader.Upgrade(ws.Response, ws.Request, nil)
			if err != nil {
				return err
			}
			client := &hub.Client{
				Connection: conn,
				Send:       make(chan []byte, 1024),
				ClientInfo: &hub.ClientInfo{
					Name:  clientInfo.Name,
					Email: clientInfo.Email,
					Token: clientInfo.Token,
				},
			}
			room.Hub.Register <- client
			room.Clients = append(room.Clients, client)
			client.Connection.WriteMessage(1, room.Hub.Buffer)
			for _, c := range room.Clients {
				log.Println(c.ClientInfo.Name)
			}
			go room.Writer(client)
			go room.Reader(client)
		}
	}
	return nil
}

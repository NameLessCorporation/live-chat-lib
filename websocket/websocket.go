package websoket

import (
	"net/http"

	"github.com/NameLessCorporation/live-chat-lib/hub"
	"github.com/NameLessCorporation/live-chat-lib/models"
	"github.com/gorilla/websocket"
)

// WebSocket ...
type WebSocket struct {
	Upgrader *websocket.Upgrader
	Hub      *hub.Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// NewWebSocket ...
func NewWebSocket(hub *hub.Hub) *WebSocket {
	return &WebSocket{
		Upgrader: &upgrader,
		Hub:      hub,
	}
}

// RunWebSocket ...
func (ws *WebSocket) RunWebSocket(wsConfig *models.WebSocketConfig, clientInfo *hub.ClientInfo) error {
	ws.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	websocket, err := ws.Upgrader.Upgrade(wsConfig.Response, &wsConfig.Request, nil)
	if err != nil {
		return err
	}
	client := &hub.Client{
		Hub:        ws.Hub,
		Connection: websocket,
		Send:       make(chan []byte, 1024),
		ClientInfo: &hub.ClientInfo{
			Name:  clientInfo.Name,
			Email: clientInfo.Email,
		},
		Token: clientInfo.Token,
	}
	room := &models.Room{
		Token: wsConfig.Token,
		Hub:   ws.Hub,
	}
	if client.Token == room.Token {
		room.Hub.Register <- client
	} else {
		websocket.WriteMessage(1, []byte("Token is incorrect"))
		websocket.Close()
	}
	if err != nil {
		return err
	}
	go client.Writer()
	go client.Reader()
	return nil
}

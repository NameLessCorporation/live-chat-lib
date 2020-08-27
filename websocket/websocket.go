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
func NewWebSocket() *WebSocket {
	return &WebSocket{
		Upgrader: &upgrader,
		Hub:      hub.NewHub(),
	}
}

// RunWebSocket ...
func (ws *WebSocket) RunWebSocket(wsConfig *models.WebSoketConfig, clientInfo *models.ClientInfo) error {
	ws.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	websocket, err := ws.Upgrader.Upgrade(wsConfig.Response, &wsConfig.Request, nil)
	if err != nil {
		return err
	}
	client := &models.Client{
		Hub:        ws.Hub,
		Connection: websocket,
		Send:       make(chan []byte, 1024),
		ClientInfo: &models.ClientInfo{
			Name:  clientInfo.Name,
			Email: clientInfo.Email,
		},
		Token: wsConfig.Token,
	}
	err = ws.Reader(websocket)
	if err != nil {
		return err
	}
	return nil
}

// Reader ...
func (ws *WebSocket) Reader(conn *websocket.Conn) error {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			return err
		}
	}
}

// Writer ...
func (ws *WebSocket) Writer(conn *websocket.Conn) {

}

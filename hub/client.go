package hub

import (
	"github.com/gorilla/websocket"
)

// Client ...
type Client struct {
	Connection *websocket.Conn
	Send       chan []byte
	ClientInfo *ClientInfo
}

// ClientInfo ...
type ClientInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

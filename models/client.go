package models

import (
	"github.com/NameLessCorporation/live-chat-lib/hub"
	"github.com/gorilla/websocket"
)

// Client ...
type Client struct {
	Hub        *hub.Hub
	Connection *websocket.Conn
	Send       chan []byte
	ClientInfo *ClientInfo
	Token      string
}

// ClientInfo ...
type ClientInfo struct {
	Name  string
	Email string
}

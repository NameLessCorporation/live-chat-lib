package hub

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 15 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024
)

// Client ...
type Client struct {
	Hub        *Hub
	Connection *websocket.Conn
	Send       chan []byte
	ClientInfo *ClientInfo
	Token      string
}

// ClientInfo ...
type ClientInfo struct {
	Name  string
	Email string
	Token string
}

// Writer ...
func (client *Client) Writer() error {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.Connection.Close()
	}()
	for {
		select {
		case message, ok := <-client.Send:
			client.Connection.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				client.Connection.WriteMessage(websocket.CloseMessage, []byte{})
			}
			w, err := client.Connection.NextWriter(websocket.TextMessage)
			if err != nil {
				return err
			}
			w.Write(message)
			n := len(client.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-client.Send)
			}
			err = w.Close()
			if err != nil {
				return err
			}
		case <-ticker.C:
			client.Connection.SetWriteDeadline(time.Now().Add(writeWait))
			err := client.Connection.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return err
			}
		}
	}
}

// Reader ...
func (client *Client) Reader() {

}

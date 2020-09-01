package models

import (
	"bytes"
	"fmt"

	"github.com/NameLessCorporation/live-chat-lib/hub"
	"github.com/gorilla/websocket"
)

const (
	maxMessageSize = 1024
)

// Room ...
type Room struct {
	Token   string
	Clients []*hub.Client
	Hub     *hub.Hub
}

// Writer ...
func (room *Room) Writer(client *hub.Client) error {
	defer func() {
		client.Connection.Close()
	}()
	for {
		data, ok := <-client.Send
		if !ok {
			client.Connection.WriteMessage(websocket.CloseMessage, []byte{})
			return nil
		}
		client.Connection.WriteMessage(1, data)
	}
}

// Reader ...
func (room *Room) Reader(client *hub.Client) error {
	defer func() {
		room.Hub.Unregister <- client
		client.Connection.Close()
	}()
	client.Connection.SetReadLimit(maxMessageSize)
	for _, c := range room.Clients {
		if client == c {
			for {
				_, mess, err := client.Connection.ReadMessage()
				if err != nil {
					return err
				}
				data := []byte(fmt.Sprintf("%s: %s\n", client.ClientInfo.Name, string(mess)))
				for _, b := range data {
					room.Hub.Buffer = append(room.Hub.Buffer, b)
				}
				data = bytes.TrimSpace(bytes.Replace(data, []byte("\n"), []byte(" "), -1))
				room.Hub.Broadcast <- data
			}
		}
	}
	return nil
}

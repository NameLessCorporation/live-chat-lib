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
	Name    string `name:"json"`
	Token   string `name:"json"`
	Clients []*hub.Client
	Hub     *hub.Hub
}

// Rooms ...
type Rooms struct {
	Rooms []*Room
}

// NewRooms ...
func NewRooms() *Rooms {
	return &Rooms{
		Rooms: nil,
	}
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

// Create ...
func (rooms *Rooms) Create(room *Room) {
	h := hub.NewHub()
	go h.Run()
	room.Hub = h
	rooms.Rooms = append(rooms.Rooms, room)
}

// Delete ...
func (rooms *Rooms) Delete(room *Room) {
	for i, r := range rooms.Rooms {
		if r.Token == room.Token {
			for _, c := range room.Clients {
				c.Connection.Close()
			}
			copy(rooms.Rooms[i:], rooms.Rooms[i+1:])
			rooms.Rooms[len(rooms.Rooms)-1] = nil
			rooms.Rooms = rooms.Rooms[:len(rooms.Rooms)-1]
		}
	}
}

package models

import "github.com/NameLessCorporation/live-chat-lib/hub"

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

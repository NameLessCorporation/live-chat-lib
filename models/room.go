package models

import (
	"github.com/NameLessCorporation/live-chat-lib/hub"
)

// Room ...
type Room struct {
	Token string
	Hub   *hub.Hub
}

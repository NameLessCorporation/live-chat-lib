package models

import "github.com/NameLessCorporation/live-chat-lib/models"

// Room ...
type Room struct {
	Token   string
	Clients []*models.Client
}

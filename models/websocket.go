package models

import (
	"net/http"
)

// WebSocketConfig ...
type WebSocketConfig struct {
	Response http.ResponseWriter
	Request  http.Request
	Token    string
}

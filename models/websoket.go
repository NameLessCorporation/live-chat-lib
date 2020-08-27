package models

import (
	"net/http"
)

// WebSoketConfig ...
type WebSoketConfig struct {
	Response http.ResponseWriter
	Request  http.Request
	Token    string
}

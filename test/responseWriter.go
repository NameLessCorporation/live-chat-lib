package main

import (
	"log"
	"net/http"
)

func NewResponseWriter(json []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(json)
	if err != nil {
		log.Println(err)
	}
}

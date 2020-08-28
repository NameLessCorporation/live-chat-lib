package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func NewRequestReader(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request error: ", err)
	}
	return body
}

package server

import (
	"log"
	"net/http"
)

func write(w http.ResponseWriter, data []byte) {
	if _, err := w.Write(data); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

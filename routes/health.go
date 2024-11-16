package routes

import (
	"log"
	"net/http"
)

// GetHealth route
func GetHealth(response http.ResponseWriter, request *http.Request) {

	log.Println("Health check request received")

	response.Header().Set("Content-Type", "application/json")
	response.Write([]byte(`{"status": "OK"}`))
}

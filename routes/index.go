package routes

import (
	"encoding/json"
	"net/http"
)

// GetIndex route (todo)
func GetIndex(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("index")
}

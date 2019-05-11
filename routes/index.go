package routes

import (
	"encoding/json"
	"net/http"
)

// Index (todo)
func getIndex(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("index")
}

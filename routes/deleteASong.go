package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saltchang/church-music-api/models"
)

// Delete A Song (todo)
func deleteSong(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, song := range models.Songs {
		if song.SID == params["sid"] {
			models.Songs = append(models.Songs[:index], models.Songs[index+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(models.Songs)
}

package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saltchang/church-music-api/models"
)

// CreateSong route (todo)
func CreateSong(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var song models.Song
	_ = json.NewDecoder(request.Body).Decode(&song)

	song.SID = strconv.Itoa(len(models.Songs) + 1)
	models.Songs = append(models.Songs, song)
	json.NewEncoder(response).Encode(song)
}

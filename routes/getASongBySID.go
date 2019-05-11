package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/saltchang/church-music-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Get A Song By SID
func getSongBySID(response http.ResponseWriter, request *http.Request) {
	// Set Header
	response.Header().Set("Content-Type", "application/json")

	// Get params from router
	params := mux.Vars(request)

	// The filter that use to find the song by SID
	filter := bson.M{"sid": params["sid"]}

	// Defined the type of result as a Song struct
	result := &models.Song{}

	// Make a context with timeout for 5s for find the expected song
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Find a song by SID (defined by the filter)
	// and decode to the result (which is a Song struct type)
	err := db.Songs.FindOne(ctx, filter).Decode(&result)
	// Catch the error if it fails
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		json.NewEncoder(response).Encode(&models.Song{})
		cancel()
		return
	}

	// If the song found, encode it to json type
	// and return the encoded result as response
	json.NewEncoder(response).Encode(&result)

	// All things done, cancel the context
	cancel()
	return

}

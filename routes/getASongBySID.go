package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/saltchang/church-music-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetSongBySID route
func GetSongBySID(response http.ResponseWriter, request *http.Request) {
	// Set Header
	response.Header().Set("Content-Type", "application/json")

	// Get params from router
	params := mux.Vars(request)

	s := strings.Split(params["s"], " ")

	// filterSlice := []bson.M{}

	// Make a context with timeout for 5s for find the expected song
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	rlist := []*mongo.SingleResult{}

	for _, sid := range s {
		filter := bson.M{"sid": sid}
		rlist = append(rlist, db.Songs.FindOne(ctx, filter))
	}

	// Defined the type of result as a Song struct
	result := []*models.Song{}

	for _, r := range rlist {
		single := &models.Song{}

		err := r.Decode(single)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(&models.Song{})
			cancel()
			return
		}

		result = append(result, single)

		// If the song found, encode it to json type
		// and return the encoded result as response

	}

	json.NewEncoder(response).Encode(&result)

	cancel()
	return

	// The filter that use to find the song by SID
	// filter := bson.M{"sid": params["sid"]}

	// Find a song by SID (defined by the filter)
	// and decode to the result (which is a Song struct type)
	// err := db.Songs.FindOne(ctx, filter).Decode(&result)
	// Catch the error if it fails

	// All things done, cancel the context

}

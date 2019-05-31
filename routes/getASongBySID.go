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

	// Split the sid params to array
	s := strings.Split(params["s"], "+")

	// Create a list of result of FindOne function
	rlist := []*mongo.SingleResult{}

	// Make a context with timeout for 5s for find the expected song
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Range the array of sid and find the result
	for _, sid := range s {
		// For each sid, make the specific filter
		filter := bson.M{"sid": sid}
		// Find the song and add it into the list of result
		rlist = append(rlist, db.Songs.FindOne(ctx, filter))
	}

	// For debugging, enable this code to see if the number of sid is correct
	// fmt.Println(len(rlist))

	// Create the result for store the array of the model of song (pointer)
	result := []*models.Song{}

	// For each searching result in rlist...
	for index, r := range rlist {
		// Get the address of a new void single song
		single := &models.Song{}

		// Decode the result to single
		err := r.Decode(single)
		if err != nil {
			// If something goes wrong(ex. song not found)...
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(bson.M{
				"error_code": 1,
				"message":    fmt.Sprintf("SID[%d]: no result found.", index),
			})
			cancel()
			return
		}

		// If everything is OK, add the found song into the result list
		result = append(result, single)

	}

	// Encode result list to json type, and return it as response
	json.NewEncoder(response).Encode(&result)

	// Done cancel the context and return
	cancel()
	return

}

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

// UpdateSong route
func UpdateSong(response http.ResponseWriter, request *http.Request) {
	// Set Header
	response.Header().Set("Content-Type", "application/json")
	// Get params from router
	params := mux.Vars(request)

	// The filter that use to find the song by SID
	filter := bson.M{"sid": params["sid"]}

	// Make a bson.M to store the data that ready to update
	newData := bson.M{}
	err := json.NewDecoder(request.Body).Decode(&newData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		json.NewEncoder(response).Encode(bson.M{
			"error_code": 5,
			"message":    "Don't play with me",
		})
		return
	}

	// Make a update interface
	update := bson.M{"$set": newData}

	// Defined the type of result as a models.Song struct
	result := &models.Song{}

	// Make a context with timeout for 5s for find the expected song
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Check if the song exist by its SID (defined by the filter)
	// and decode to the result (which is a Song struct type)
	err = db.Songs.FindOne(ctx, filter).Decode(&result)
	// Catch the error if it fails
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		json.NewEncoder(response).Encode(bson.M{
			"error_code": 4,
			"message":    "No result found.",
		})
		cancel()
		return
	}

	// Update the song
	res, err := db.Songs.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		json.NewEncoder(response).Encode(bson.M{
			"error_code": 3,
			"message":    "Something Wrong",
		})
		cancel()
		return
	}

	// If the song found, encode it to json type
	// and return the encoded result as response
	json.NewEncoder(response).Encode(&res)

	// All things done, cancel the context
	cancel()
	return
}

package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/saltchang/church-music-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteSong route (todo)
func DeleteSong(response http.ResponseWriter, request *http.Request) {
	// Set Header
	response.Header().Set("Content-Type", "application/json")
	// Get params from router
	params := mux.Vars(request)

	// Make a Token sturct to store the token data
	var token models.Token

	// Read the body
	body, err := ioutil.ReadAll(request.Body)

	// decode the body as token struct
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		json.NewEncoder(response).Encode(bson.M{
			"Code":    1200,
			"Message": "Don't play with me",
		})
		return
	}

	// Make a context with timeout for 5s for find the expected song
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Token filter that used to find the token in db
	tokenFilter := bson.M{"token": token.Token}

	// Check the token
	err = db.Tokens.FindOne(ctx, tokenFilter).Decode(&token)
	// Catch the error if it fails
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		json.NewEncoder(response).Encode(bson.M{
			"Code":    1110,
			"Message": "Not correct access token",
		})
		cancel()
		return
	}

	// Check if the level of token is the most high.
	if token.Autho != "MOSTHIGHADMIN" {
		fmt.Printf("Error: %v\n", err)
		json.NewEncoder(response).Encode(bson.M{
			"Code":    1120,
			"Message": "This token has no authority",
		})
		cancel()
		return
	}

	// The filter that used to find the song by SID
	filter := bson.M{"sid": params["sid"]}

	// Defined the type of result as a models.Song struct
	result := &models.Song{}

	// Check if the song exist by its SID and delete it
	err = db.Songs.FindOneAndDelete(ctx, filter).Decode(&result)
	// Catch the error if it fails
	if err != nil {
		fmt.Printf("Error deleted song, sid: %v\n", params["sid"])
		fmt.Printf("Error: %v\n", err)
		json.NewEncoder(response).Encode(bson.M{
			"Code":      1500,
			"DeleteSID": params["sid"],
			"Message":   "Delete song error.",
		})
		cancel()
		return
	}

	// Response
	json.NewEncoder(response).Encode(bson.M{
		"Code":      1000,
		"DeleteSID": params["sid"],
		"Message":   "Delete song Successfully",
	})

	// If the song deleted successfully, encode it to json type
	// and return the encoded result as response
	// json.NewEncoder(response).Encode(&result)

	// All things done, cancel the context
	cancel()
	return
}

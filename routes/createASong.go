package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/saltchang/church-music-api/helper"
	"github.com/saltchang/church-music-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	sidGenerator = helper.SidGenerator
)

// CreateSong route
func CreateSong(response http.ResponseWriter, request *http.Request) {
	// Set Header
	response.Header().Set("Content-Type", "application/json")

	// Make a Token sturct to store the token data
	var token models.Token

	// Read the body
	body, err := ioutil.ReadAll(request.Body)

	// decode the body as token struct
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		json.NewEncoder(response).Encode(bson.M{
			"Code":    1100,
			"Message": "Broken during render token.",
		})
		return
	}

	// Make a bson.M to store the data that ready to update
	newData := bson.M{}

	// err = json.NewDecoder(request.Body).Decode(&newData)
	err = json.Unmarshal(body, &newData)
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

	// Delete the token column
	delete(newData, "token")

	newSongLang := fmt.Sprintf("%v", newData["language"])
	newSongNumC := fmt.Sprintf("%v", newData["num_c"])
	newSongNumI := fmt.Sprintf("%v", newData["num_i"])

	// Create new sid by song data
	success, newSID := sidGenerator.GenSID(newSongLang, newSongNumC, newSongNumI)

	if !success {
		fmt.Println("Error song information.")
		json.NewEncoder(response).Encode(bson.M{
			"Code":    1300,
			"Message": "Error song information.",
		})
		cancel()
		return
	}

	songFilter := bson.M{"sid": newSID}
	var checkResult models.Song
	// Find the song by sid
	err = db.Songs.FindOne(ctx, songFilter).Decode(&checkResult)
	// Check is the song exist
	if err == nil {
		fmt.Println("Error: sid exist.")
		json.NewEncoder(response).Encode(bson.M{
			"Code":    1400,
			"Message": "Song sid exist.",
		})
		cancel()
		return
	}

	newData["sid"] = newSID

	res, err := db.Songs.InsertOne(ctx, newData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		json.NewEncoder(response).Encode(bson.M{
			"Code":    1410,
			"Message": "Creating song is fail.",
		})
		cancel()
		return
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)

	result := bson.M{
		"Code":      1000,
		"Message":   "Create a song data successfully.",
		"NewSID":    newSID,
		"NewSongID": res.InsertedID,
	}

	// Update the song

	// If the song found, encode it to json type
	// and return the encoded result as response
	json.NewEncoder(response).Encode(result)

	// All things done, cancel the context
	cancel()
	return
}

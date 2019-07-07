package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetSongBySearch route
func GetSongBySearch(response http.ResponseWriter, request *http.Request) {
	// Set Header
	response.Header().Set("Content-Type", "application/json")
	// Get params from router
	params := mux.Vars(request)

	// Decode the params as args

	// If Testing mode
	TestingMode := false
	if params["test"] == "1" {
		TestingMode = true
	}

	// Song.Language arg
	lang := strings.Replace(params["lang"], " ", "", -1)
	// Song.NumC arg
	numc := strings.Replace(params["c"], " ", "", -1)
	// Song.Tonality arg
	tona := strings.Replace(params["to"], " ", "", -1)
	// Song.Title arg
	titleQ := params["title"]
	// Split the arg by space, it was displayed as "+"
	titles := strings.Split(titleQ, " ")
	// Song.Lyrics arg
	lyricsQ := params["lyrics"]
	// Split the arg by space, it was displayed as "+"
	lyrics := strings.Split(lyricsQ, " ")

	// Make a slice of bson.M prepared to jonin in the filter
	filterSlice := []bson.M{}

	// If there's a lang arg, add the lang filter into the filter slice
	if lang != "" {
		filterSlice = append(filterSlice, bson.M{"language": lang})
	}
	// If there's a num_c arg, add the lang filter into the filter slice
	if numc != "" {
		filterSlice = append(filterSlice, bson.M{"num_c": numc})
	}
	// If there's a num_c arg, add the lang filter into the filter slice
	if tona != "" {
		filterSlice = append(filterSlice, bson.M{"tonality": tona})
	}

	// Add all the keyword args of title into the filter slice
	for _, s := range titles {
		// If the title arg is not empty
		if s != "" {
			filterSlice = append(filterSlice, bson.M{"title": bson.M{"$regex": s}})
		}
	}

	// Add all the keyword args of lyrics into the filter slice
	for _, lyric := range lyrics {
		// If the title arg is not empty
		if lyric != "" {
			filterSlice = append(filterSlice, bson.M{"lyrics": bson.M{"$regex": lyric}})
		}
	}

	// If she has no any key, then she shall be saved by me, Hero of the World.
	if len(filterSlice) == 0 {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(bson.M{
			"error_code": 1,
			"message":    "Don't play with me",
		})
		return
	}

	// Make the filter and put all conditions from slice into it
	filter := bson.M{"$and": filterSlice}

	// Make a options for sorting the songs result
	opts := options.FindOptions{Sort: bson.M{"sid": 1}}

	// Make a context with timeout for 30s, for listing songs
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// Create a cursor to search the songs by all args
	var cur *mongo.Cursor
	var err error
	if TestingMode {
		cur, err = db.SongsForTesting.Find(ctx, filter, &opts)
	} else {
		cur, err = db.Songs.Find(ctx, filter, &opts)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		cancel()
		return
	}
	// Make a defer to handle the closing of the cursor
	defer cur.Close(ctx)

	// Make a slice for saving the result songs
	var list []bson.M

	// If there's next song in the cursor
	for cur.Next(ctx) {
		// Make a bson.M type result buffer
		var result bson.M

		// Decode the current song pointed by the cursor as result
		err := cur.Decode(&result)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			cancel()
			return
		}

		// Save the current found song into the slice
		list = append(list, result)
	}

	// If there's any error in the cursor
	if err := cur.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
		cancel()
		return
	}

	if len(list) == 0 {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(bson.M{
			"error_code": 2,
			"message":    "No result found.",
		})
		cancel()
		return
	}

	// All songs found out, Encode the songs list to json and return it as a response
	json.NewEncoder(response).Encode(&list)

	cancel()
	return

}

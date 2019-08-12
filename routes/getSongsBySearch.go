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
	// Fuzzy search preprocessor
	titleQList := []string{titleQ}
	// Song.Lyrics arg
	lyricsQ := params["lyrics"]
	// Split the arg by space, it was displayed as "+"
	// Fuzzy search preprocessor
	lyricsQList := []string{lyricsQ}

	// Check titles content for "你, 祢, 他, 祂"
	if strings.Contains(titleQ, "你") {
		titleQList = append(titleQList, strings.Replace(titleQ, "你", "祢", -1))
	}
	if strings.Contains(titleQ, "祢") {
		titleQList = append(titleQList, strings.Replace(titleQ, "祢", "你", -1))
	}
	if strings.Contains(titleQ, "他") {
		titleQList = append(titleQList, strings.Replace(titleQ, "他", "祂", -1))
	}
	if strings.Contains(titleQ, "祂") {
		titleQList = append(titleQList, strings.Replace(titleQ, "祂", "他", -1))
	}

	// Check lyrics content for "你, 祢, 他, 祂"
	if strings.Contains(lyricsQ, "你") {
		lyricsQList = append(lyricsQList, strings.Replace(lyricsQ, "你", "祢", -1))
	}
	if strings.Contains(lyricsQ, "祢") {
		lyricsQList = append(lyricsQList, strings.Replace(lyricsQ, "祢", "你", -1))
	}
	if strings.Contains(lyricsQ, "他") {
		lyricsQList = append(lyricsQList, strings.Replace(lyricsQ, "他", "祂", -1))
	}
	if strings.Contains(lyricsQ, "祂") {
		lyricsQList = append(lyricsQList, strings.Replace(lyricsQ, "祂", "他", -1))
	}

	// list of splited titles: ex. [ ["你", "愛"], ["祢", "愛"] ]
	titleSplitedList := [][]string{}
	for _, eachTitleQ := range titleQList {
		titleSplitedList = append(titleSplitedList, strings.Split(eachTitleQ, " "))
	}

	// Fuzzy lyrics slice
	lyricsSplitedList := [][]string{}
	for _, eachLyricsQ := range lyricsQList {
		lyricsSplitedList = append(lyricsSplitedList, strings.Split(eachLyricsQ, " "))
	}

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

	// Make titles filter fuzzyQuerySlice in bson.A type
	eachTitleQQuerySlice := bson.A{}
	// Get titles in titleSplitedList
	for _, titles := range titleSplitedList {
		// Add all the keyword args of title into the filter slice
		tempTitleFilter := bson.A{}
		for _, s := range titles {
			// If the title arg is not empty
			if s != "" {
				tempTitleFilter = append(tempTitleFilter, bson.M{"title": bson.M{"$regex": s}})
			}
		}
		if len(tempTitleFilter) > 0 {
			eachTitleQQuerySlice = append(eachTitleQQuerySlice, bson.M{"$and": tempTitleFilter})
		}
	}
	if len(eachTitleQQuerySlice) > 0 {
		fmt.Println(eachTitleQQuerySlice)
		// Make a title slice for title query Slice
		titleFilter := bson.M{"$or": eachTitleQQuerySlice}
		filterSlice = append(filterSlice, titleFilter)
	}

	// Make lyrics filter fuzzyQuerySlice in bson.A type
	eachLyricsQQuerySlice := bson.A{}
	// Get lyrics in lyricsSplitedList
	for _, lyrics := range lyricsSplitedList {
		// Add all the keyword args of lyrics into the filter slice
		tempLyricsFilter := bson.A{}
		for _, s := range lyrics {
			// If the lyrics arg is not empty
			if s != "" {
				tempLyricsFilter = append(tempLyricsFilter, bson.M{"lyrics": bson.M{"$regex": s}})
			}
		}
		if len(tempLyricsFilter) > 0 {
			eachLyricsQQuerySlice = append(eachLyricsQQuerySlice, bson.M{"$and": tempLyricsFilter})
		}
	}
	if len(eachLyricsQQuerySlice) > 0 {
		// Make a lyrics slice for lyrics query Slice
		lyricsFilter := bson.M{"$or": eachLyricsQQuerySlice}
		filterSlice = append(filterSlice, lyricsFilter)
	}
	fmt.Println(filterSlice)

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

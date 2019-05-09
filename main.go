package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	songs   []Song            // Songs data model
	songsDB *mongo.Collection // The collection of songs data in MongoDB
)

// Song Struct: This is the songs data model.
type Song struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SID           string             `json:"sid" bson:"sid"`
	NumC          string             `json:"num_c" bson:"num_c"`
	NumI          string             `json:"num_i" bson:"num_i"`
	Title         string             `json:"title" bson:"title"`
	Album         string             `json:"album" bson:"album"`
	Tonality      string             `json:"tonality" bson:"tonality"`
	Year          string             `json:"year" bson:"year"`
	Language      string             `json:"language" bson:"language"`
	Lyrics        [][]string         `json:"lyrics" bson:"lyrics"`
	Tempo         string             `json:"tempo" bson:"tempo"`
	TimeSignature string             `json:"time_signature" bson:"time_signature"`
	Publisher     string             `json:"publisher" bson:"publisher"`
	Lyricist      string             `json:"lyricist" bson:"lyricist"`
	Composer      string             `json:"composer" bson:"composer"`
	Translator    string             `json:"translator" bson:"translator"`
}

// Dummy Data: This is for development.
func dummySongs() {
	songs = append(songs,
		Song{
			SID:      "1011054",
			NumC:     "11",
			NumI:     "54",
			Title:    "我獻上我心",
			Album:    "這是真愛",
			Tonality: "G",
			Year:     "",
			Language: "Chinese",
			Lyrics: [][]string{
				[]string{
					"p",
					"我心何等渴望，來尊崇你，主，我用全心來敬拜你，",
					"凡在我裡面的，都讚美你，我一切所愛，在於你。",
				},
				[]string{
					"p",
					"主，我獻上我心，我獻上我的靈，",
					"我活著為了你，我的每個氣息，",
					"生命中的每個時刻，主，成全你旨意。",
				},
				[]string{
					"p",
					"獻上我心，獻上我靈。",
				},
			},
			Tempo:         "",
			TimeSignature: "",
			Publisher:     "",
			Lyricist:      "Reuben Morgan",
			Composer:      "Reuben Morgan",
			Translator:    "周巽光",
		})

	songs = append(songs,
		Song{
			SID:      "1010066",
			NumC:     "10",
			NumI:     "66",
			Title:    "前來敬拜",
			Album:    "新的事將要成就",
			Tonality: "G",
			Year:     "",
			Language: "Chinese",
			Lyrics: [][]string{
				[]string{
					"v",
					"哈利路亞，哈利路亞，前來敬拜永遠的君王，",
					"哈利路亞，哈利路亞，大聲宣告主榮耀降臨。",
				},
				[]string{
					"c",
					"榮耀尊貴，能力權柄歸於你，",
					"你是我的救主，我的救贖，",
					"榮耀尊貴，能力權柄歸於你，",
					"你是配得，你是配得，你是配得我的敬拜。",
				},
				[]string{
					"b",
					"榮耀尊貴，美麗無比，神的兒子，耶穌我的主，",
					"榮耀尊貴，美麗無比，神的兒子，耶穌我的主。",
				},
			},
			Tempo:         "",
			TimeSignature: "",
			Publisher:     "",
			Lyricist:      "Reuben Morgan",
			Composer:      "Reuben Morgan",
			Translator:    "周巽光",
		})
}

// Index (todo)
func getIndex(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("index")
}

// Get All Songs
func getAllSongs(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Make a context with timeout for 30s, for listing all songs
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// Make a options for sorting the songs result
	opts := options.FindOptions{Sort: bson.M{"sid": 1}}
	// Create a cursor to find the songs
	cur, err := songsDB.Find(ctx, bson.D{}, &opts)
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

	// All songs found out, Encode the songs list to json and return it as a response
	json.NewEncoder(response).Encode(&list)

	cancel()
	return

}

// Get A Song By SID
func getSongBySID(response http.ResponseWriter, request *http.Request) {
	// Set Header
	response.Header().Set("Content-Type", "application/json")

	// Get params from router
	params := mux.Vars(request)

	// The filter that use to find the song by SID
	filter := bson.M{"sid": params["sid"]}

	// Defined the type of result as a Song struct
	result := &Song{}

	// Make a context with timeout for 5s for find the expected song
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Find a song by SID (defined by the filter)
	// and decode to the result (which is a Song struct type)
	err := songsDB.FindOne(ctx, filter).Decode(&result)
	// Catch the error if it fails
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		json.NewEncoder(response).Encode(&Song{})
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

// Get A Song By Multiple Queries Searching
func getSongBySearch(response http.ResponseWriter, request *http.Request) {
	// Set Header
	response.Header().Set("Content-Type", "application/json")
	// Get params from router
	params := mux.Vars(request)

	// Decode the params as args
	// Song.Language arg
	lang := params["lang"]
	// Song.NumC arg
	numc := params["c"]
	// Song.Tonality arg
	tona := params["to"]
	// Song.Title arg
	titleQ := params["title"]
	// Split the arg by space, it was displayed as "+"
	titles := strings.Split(titleQ, " ")

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
			filterSlice = append(filterSlice, bson.M{"title": primitive.Regex{Pattern: s, Options: ""}})
		}
	}

	// If she has no any key, then she shall be saved by me, Hero of the World.
	if len(filterSlice) == 0 {
		json.NewEncoder(response).Encode(bson.M{
			"error_code": 5,
			"message":    "Don't play with me",
		})
		return
	}

	// Make the filter and put all conditions from slice into it
	filter := bson.M{"$and": filterSlice}

	// Make a context with timeout for 30s, for listing songs
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// Create a cursor to search the songs by all args
	cur, err := songsDB.Find(ctx, filter)
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
		json.NewEncoder(response).Encode(bson.M{
			"error_code": 4,
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

// Create A Song (todo)
func createSong(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var song Song
	_ = json.NewDecoder(request.Body).Decode(&song)

	song.SID = strconv.Itoa(len(songs) + 1)
	songs = append(songs, song)
	json.NewEncoder(response).Encode(song)
}

// Update A Song
func updateSong(response http.ResponseWriter, request *http.Request) {
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

	// Defined the type of result as a Song struct
	result := &Song{}

	// Make a context with timeout for 5s for find the expected song
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Check if the song exist by its SID (defined by the filter)
	// and decode to the result (which is a Song struct type)
	err = songsDB.FindOne(ctx, filter).Decode(&result)
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
	res, err := songsDB.UpdateOne(ctx, filter, update)
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

// Delete A Song (todo)
func deleteSong(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, song := range songs {
		if song.SID == params["sid"] {
			songs = append(songs[:index], songs[index+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(songs)
}

func main() {

	// MongoDB
	fmt.Print("Create Client and connected to MongoDB...")

	// Make a context with timeout for 10s for create the client for MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Create the client at port:27017 (MongoDB default)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// If it fails...
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	// Make a context with timeout for 2s for connect to MongoDB
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	// Try to connect to MongoDB and catch the error if it fails
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	// If success
	fmt.Println("...[Success!]")

	// Get MongoDB collection "songs" from database "caten-worship" as a
	// *mongo.Collection type
	songsDB = client.Database("caten-worship").Collection("songs")

	// Create the dummy songs data, just for development
	dummySongs()

	// Set the Main Router
	mainRouter := mux.NewRouter()

	// Route Handlers and Endpoints

	// Route: Home
	mainRouter.HandleFunc("/", getIndex).Methods("GET")

	// Route: Get all songs
	mainRouter.HandleFunc("/api/songs", getAllSongs).Methods("GET")

	// Route: Get a song by its SID
	mainRouter.HandleFunc("/api/songs/sid/{sid}", getSongBySID).Methods("GET")

	// Route: Get a song by searching title
	mainRouter.HandleFunc("/api/songs/search/", getSongBySearch).Queries("title", "{title}", "lang", "{lang}", "c", "{c}", "to", "{to}").Methods("GET")

	// Route: Create a song
	mainRouter.HandleFunc("/api/songs", createSong).Methods("POST")

	// Route: Update a song by its SID
	mainRouter.HandleFunc("/api/songs/sid/{sid}", updateSong).Methods("PUT")

	// Route: Delete a song by its SID
	mainRouter.HandleFunc("/api/songs/{sid}", deleteSong).Methods("DELETE")

	// All things are good now, server starts to run
	fmt.Println("Server starts to run at: http://localhost:7700")
	log.Fatal(http.ListenAndServe(":7700", mainRouter))
}

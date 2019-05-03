package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	songs []Song
)

// Song Struct (MODEL)
type Song struct {
	ID       string     `json:"id"`
	SID      *SID       `json:"sid"`
	Title    string     `json:"title"`
	Album    string     `json:"album"`
	Language string     `json:"language"`
	Lyrics   [][]string `json:"lyrics"`
	Author   *Author    `json:"authors"`
}

// SID Struct
type SID struct {
	U string `json:"u"`
	C string `json:"c"`
	I string `json:"i"`
}

// Author struct
type Author struct {
	Lyricist   string `json:"lyricist"`
	Composer   string `json:"composer"`
	Translator string `json:"translator"`
}

// Dummy Data
func dummySongs() {
	songs = append(songs,
		Song{
			ID:       "1",
			SID:      &SID{U: "1011054", C: "11", I: "54"},
			Title:    "我獻上我心",
			Album:    "這是真愛",
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
			Author: &Author{
				Lyricist:   "Reuben Morgan",
				Composer:   "Reuben Morgan",
				Translator: "周巽光"}})

	songs = append(songs,
		Song{
			ID:       "2",
			SID:      &SID{U: "1010066", C: "10", I: "66"},
			Title:    "前來敬拜",
			Album:    "新的事將要成就",
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
			Author: &Author{
				Lyricist:   "鄭懋柔",
				Composer:   "游智婷",
				Translator: ""}})
}

// Index
func getIndex(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("index")
}

// Get All Songs
func getSongs(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(songs)
}

// Get A Song
func getSong(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, song := range songs {
		if song.ID == params["id"] {
			json.NewEncoder(response).Encode(song)
			return
		}
	}
	json.NewEncoder(response).Encode(&Song{})
}

// Create A Song
func createSong(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var song Song
	_ = json.NewDecoder(request.Body).Decode(&song)

	song.ID = strconv.Itoa(len(songs) + 1)
	songs = append(songs, song)
	json.NewEncoder(response).Encode(song)
}

// Update A Song
func updateSong(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, song := range songs {
		if song.ID == params["id"] {
			songs = append(songs[:index], songs[index+1:]...)
			var updateSong Song
			_ = json.NewDecoder(request.Body).Decode(&updateSong)

			updateSong.ID = params["id"]
			songs = append(songs, updateSong)
			json.NewEncoder(response).Encode(updateSong)
			return
		}
	}
	json.NewEncoder(response).Encode(songs)
}

// Delete A Song
func deleteSong(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, song := range songs {
		if song.ID == params["id"] {
			songs = append(songs[:index], songs[index+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(songs)
}

func main() {

	// MongoDB
	// Create the MongoDB client
	fmt.Print("Connected to MongoDB...")
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	// Check the MongoDB connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("...[Success!]")

	// Get MongoDB collection "songs" from database "caten-worship"
	songsDB := client.Database("caten_songs").Collection("songs")
	fmt.Println(songsDB)

	// Songs Data
	fmt.Print("Create Dummy Songs Data...")
	dummySongs()
	fmt.Println("...[Success!]")

	// Set the Main Router
	mainRouter := mux.NewRouter()

	// Route Handlers / Endpoint
	mainRouter.HandleFunc("/", getIndex).Methods("GET")
	mainRouter.HandleFunc("/api/songs", getSongs).Methods("GET")
	mainRouter.HandleFunc("/api/songs/{id}", getSong).Methods("GET")
	mainRouter.HandleFunc("/api/songs", createSong).Methods("POST")
	mainRouter.HandleFunc("/api/songs/{id}", updateSong).Methods("PUT")
	mainRouter.HandleFunc("/api/songs/{id}", deleteSong).Methods("DELETE")
	fmt.Println("Server starts to run...")
	log.Fatal(http.ListenAndServe(":8000", mainRouter))
}

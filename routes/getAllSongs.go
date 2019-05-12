package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllSongs route
func GetAllSongs(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Make a context with timeout for 30s, for listing all songs
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// Make a options for sorting the songs result
	opts := options.FindOptions{Sort: bson.M{"sid": 1}}
	// Create a cursor to find the songs
	cur, err := db.Songs.Find(ctx, bson.D{}, &opts)
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
	fmt.Println(`"/api/songs" route accessed.`)
	json.NewEncoder(response).Encode(&list)

	cancel()
	return

}

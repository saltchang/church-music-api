package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Database struct
type Database struct {
	Songs *mongo.Collection // The collection of songs data in MongoDB
}

// InitDB function
func (db *Database) InitDB() *Database {

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
	db.Songs = client.Database("caten-worship").Collection("songs")

	cancel()

	return db
}

var (
	// DB var
	DB = new(Database)
)

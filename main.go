package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/saltchang/church-music-api/env"
	"github.com/saltchang/church-music-api/models"
	"github.com/saltchang/church-music-api/routes"
)

var (
	r  *routes.Routers
	db = models.DB
)

func main() {

	db.InitDB()

	// Create the dummy songs data, just for development
	models.DummySongs()

	// Set the Main Router
	mainRouter := r.InitRouters()

	// All things are good now, server starts to run
	fmt.Println("Server starts to run at port" + env.ENV.Port)
	log.Fatal(http.ListenAndServe(env.ENV.Port, mainRouter))
}

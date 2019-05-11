package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/saltchang/church-music-api/models"
	"github.com/saltchang/church-music-api/routes"
)

var (
	r         *routes.Routers
	db        = models.DB
	servePORT = ":7700" // Defined the PORT to serve
)

func main() {

	db.InitDB()

	// Create the dummy songs data, just for development
	models.DummySongs()

	// Set the Main Router
	mainRouter := r.InitRouters()

	// All things are good now, server starts to run
	fmt.Println("Server starts to run at PORT" + servePORT)
	log.Fatal(http.ListenAndServe(servePORT, mainRouter))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/saltchang/church-music-api/env"
	"github.com/saltchang/church-music-api/jobs"
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

	// Initialize and start the health checker
	healthChecker := jobs.NewHealthChecker(
		env.ENV.MusicAppURL+"/api/health",
		5*time.Minute,
	)
	healthChecker.Start()

	// Set the Main Router
	mainRouter := r.InitRouters()

	// All things are good now, server starts to run
	fmt.Println("Environment: ", env.ENV.AppConfig)
	fmt.Println("Server starts to run at port: " + env.ENV.Port)
	log.Fatal(http.ListenAndServe(":"+env.ENV.Port, mainRouter))
}

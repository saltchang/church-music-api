package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// ENV global
	ENV = new(Env).loadENV()
)

// Env struct
type Env struct {
	TestVar             string
	Port                string
	MongoURI            string
	SongsDBName         string
	SongsCollectionName string
}

func (env *Env) loadENV() *Env {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/saltchang/church-music-api/.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env.TestVar = os.Getenv("TEST_VAR")
	env.Port = os.Getenv("PORT")
	env.MongoURI = os.Getenv("MONGO_URI")
	env.SongsDBName = os.Getenv("SONGS_DB_NAME")
	env.SongsCollectionName = os.Getenv("SONGS_COLLECTION_NAME")

	return env
}

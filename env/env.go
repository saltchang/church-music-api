package env

import (
	"fmt"
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
	AppConfig            string
	TestVar              string
	Port                 string
	MongoURI             string
	SongsDBName          string
	SongsCollectionName  string
	TokensCollectionName string
}

func (env *Env) loadENV() *Env {
	if os.Getenv("APP_CONFIG") != "PRODUCTION" {
		err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/saltchang/church-music-api/.env"))
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	env.AppConfig = os.Getenv("APP_CONFIG")
	env.TestVar = os.Getenv("TEST_VAR")
	env.Port = os.Getenv("PORT")
	env.MongoURI = os.Getenv("MONGO_URI")
	env.SongsDBName = os.Getenv("SONGS_DB_NAME")
	env.SongsCollectionName = os.Getenv("SONGS_COLLECTION_NAME")
	env.TokensCollectionName = os.Getenv("TOKENS_COLLECTION_NAME")

	fmt.Println("Port set:", env.Port)

	return env
}

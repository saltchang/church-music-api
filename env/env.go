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
	AppConfig                 string
	AppEnv                    string
	EnvFile                   string
	TestVar                   string
	Port                      string
	MongoURI                  string
	SongsDBName               string
	SongsCollectionName       string
	SongsCollectionNameForDev string
	TokensCollectionName      string
	MusicAppURL               string
}

func (env *Env) loadENV() *Env {
	if os.Getenv("APP_CONFIG") != "PRODUCTION" {
		env.EnvFile = os.Getenv("ENV_FILE")
		fmt.Println("EnvFile:", env.EnvFile)

		err := godotenv.Load(os.ExpandEnv(env.EnvFile))
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	env.AppEnv = os.Getenv("APP_ENV")
	env.AppConfig = os.Getenv("APP_CONFIG")
	env.TestVar = os.Getenv("TEST_VAR")
	env.Port = os.Getenv("PORT")
	env.MongoURI = os.Getenv("MONGO_URI")
	env.SongsDBName = os.Getenv("SONGS_DB_NAME")
	env.SongsCollectionName = os.Getenv("SONGS_COLLECTION_NAME")
	env.SongsCollectionNameForDev = os.Getenv("SONGS_COLLECTION_NAME_FOR_TEST")
	env.TokensCollectionName = os.Getenv("TOKENS_COLLECTION_NAME")
	env.MusicAppURL = os.Getenv("MUSIC_APP_URL")

	fmt.Println("Environment:", env.AppEnv)
	fmt.Println("Port set:", env.Port)

	return env
}

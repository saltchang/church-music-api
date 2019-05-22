package routes

import (
	"net/http"
)

// GetIndex route (todo)
func GetIndex(response http.ResponseWriter, request *http.Request) {

	response.Write([]byte("<h1>Church Music API</h1>\n<p>By Salt</p>\n<a href='https://github.com/saltchang/church-music-api'>GitHub</a>"))

	return
}

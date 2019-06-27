package routes

import (
	"net/http"
)

// GetFavicon route
func GetFavicon(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "favicon.ico")
	return
}

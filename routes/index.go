package routes

import (
	"encoding/json"
	"net/http"
)

type PageData struct {
	PageTitle string
	Todos     []Todo
}

type Todo struct {
	Title string
	Done  bool
}

// GetIndex route (todo)
func GetIndex(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("index")
	// data := PageData{
	// 	PageTitle: "My TODO list",
	// 	Todos: []Todo{
	// 		{Title: "Task 1", Done: false},
	// 		{Title: "Task 2", Done: true},
	// 		{Title: "Task 3", Done: true},
	// 	},
	// }

	// tmpl, err := template.ParseFiles("templates/index.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// tmpl.Execute(response, PageData{})
	return
}

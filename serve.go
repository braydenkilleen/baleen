package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/braydenkilleen/baleen/database"
	"github.com/braydenkilleen/baleen/models"
	"github.com/gorilla/mux"
)

// // PageData ...
type PageData struct {
	Items []models.Item
}

// type spaHandler struct {
// 	staticPath string
// 	indexPath  string
// }

// ServeHTTP ...
// func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// }

// Serve starts a webserver
func Serve() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		items, err := database.AllItems()
		if err != nil {
			log.Fatal(err)
		}

		// data := PageData{
		// 	Bookmarks: bookmarks,
		// 	Page:      "bookmarks",
		// }
		data := PageData{
			Items: items,
		}
		tmpl := template.Must(template.ParseFiles("web/templates/layout.html"))
		tmpl.Execute(w, data)
	})

	log.Println("listening on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}

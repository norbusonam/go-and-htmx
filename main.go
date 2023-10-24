package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Database struct {
	Posts []Post
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := Database{}

	// parse templates
	tmpl := template.Must(template.ParseFiles("index.html"))

	// handle root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		posts := db.Posts
		tmpl.Execute(w, posts)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

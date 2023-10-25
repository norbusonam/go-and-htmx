package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Post struct {
	Id    int
	Title string
	Body  string
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
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	// handle root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, db.Posts)
	})

	// create post
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			db.Posts = append(db.Posts, Post{
				Id:    len(db.Posts) + 1,
				Title: r.FormValue("title"),
				Body:  r.FormValue("body"),
			})
		}
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

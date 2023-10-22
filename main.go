package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
		if err != nil {
			log.Fatal(err)
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		posts := []Post{}
		err = json.Unmarshal(b, &posts)
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, posts)
	}

	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

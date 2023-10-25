package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type ViewData struct {
	Counter int
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ==================== COUNTER ====================
	counter := 0
	http.HandleFunc("/counter/increment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			counter++
			w.Write([]byte(fmt.Sprintf("%d", counter)))
		}
	})

	http.HandleFunc("/counter/decrement", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			counter--
			w.Write([]byte(fmt.Sprintf("%d", counter)))
		}
	})

	// ==================== INDEX ====================
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, ViewData{
			Counter: counter,
		})
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

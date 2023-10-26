package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ViewData struct {
	Counter   int
	ListItems []int
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	// ==================== COUNTER ====================
	counter := 0
	http.HandleFunc("/counter/increment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			counter++
			w.Write([]byte(strconv.Itoa(counter)))
		}
	})

	http.HandleFunc("/counter/decrement", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			counter--
			w.Write([]byte(strconv.Itoa(counter)))
		}
	})

	// ==================== LIST CONTROLS ====================
	listItems := []int{}
	nextListItem := 1
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			listItems = append(listItems, nextListItem)
			nextListItem++
			tmpl.ExecuteTemplate(w, "list", ViewData{
				ListItems: listItems,
			})
		}
	})

	// ==================== INDEX ====================
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, ViewData{
			Counter:   counter,
			ListItems: listItems,
		})
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

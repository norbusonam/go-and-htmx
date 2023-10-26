package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ViewData struct {
	Counter       int
	ListItems     []int
	SearchResults []string
}

func main() {
	// set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// parse template
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	// =================================================
	// ==================== COUNTER ====================
	// =================================================

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

	// =======================================================
	// ==================== LIST CONTROLS ====================
	// =======================================================

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

	http.HandleFunc("/list/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			itemToDelete, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
			if err != nil {
				log.Fatal(err)
			}
			for i, item := range listItems {
				if item == itemToDelete {
					listItems = append(listItems[:i], listItems[i+1:]...)
					break
				}
			}
			tmpl.ExecuteTemplate(w, "list", ViewData{
				ListItems: listItems,
			})
		}
	})

	http.HandleFunc("/list/reset", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			listItems = []int{}
			nextListItem = 1
			tmpl.ExecuteTemplate(w, "list", ViewData{
				ListItems: listItems,
			})
		}
	})

	// =======================================================
	// ==================== ACTIVE SEARCH ====================
	// =======================================================

	names := []string{"Norbu", "Ryan", "Ben", "Adam", "Christian", "Bea"}

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			query := r.URL.Query().Get("query")
			matches := []string{}
			if query != "" {
				for _, name := range names {
					if strings.Contains(strings.ToLower(name), strings.ToLower(query)) {
						matches = append(matches, name)
					}
				}
			}
			tmpl.ExecuteTemplate(w, "search-results", ViewData{
				SearchResults: matches,
			})
		}
	})

	// ===============================================
	// ==================== INDEX ====================
	// ===============================================

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, ViewData{
			Counter:   counter,
			ListItems: listItems,
		})
	})

	// start server
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

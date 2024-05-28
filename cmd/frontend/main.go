package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	
	// Register the handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w)
	})
	
	// Start the server
	log.Println("Server started on: http://localhost:80")
	// the variable err is assigned and returned at the same time
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter) {
	log.Println("Rendering template...")

	partials := []string{
		// content.html should be the first file
		"./cmd/frontend/templates/content.html",
		"./cmd/frontend/templates/layout.html",
		"./cmd/frontend/templates/head.html",
	}

	tmpl, err := template.ParseFiles(partials...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// execute the template and return Internal Server Error if it fails
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



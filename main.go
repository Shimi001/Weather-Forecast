package main

import (
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}

	// Set up a simple web server
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Weather API endpoint
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			http.Error(w, "API_KEY not set in .env file", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":80", nil)
}

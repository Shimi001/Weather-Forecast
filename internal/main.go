package main

import (
	"net/http"
	"text/template"

	"github.com/joho/godotenv"
)

// Weather API response structure
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}

	setApiKey()

	// Set up a simple web server
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve the main page
	http.HandleFunc("/", mainPageHandler(tmpl))
	http.HandleFunc("/weather", weatherHandler())

	// Start the server on port 80
	http.ListenAndServe(":80", nil)
}

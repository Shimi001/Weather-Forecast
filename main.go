package main

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

type WeatherData struct {
	City string `json:"city"`
}

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

		сity := r.URL.Query().Get("city")
		if сity == "" {
			http.Error(w, "City parameter is required", http.StatusBadRequest)
			return
		}

		data := WeatherData{
			City: сity,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":80", nil)
}

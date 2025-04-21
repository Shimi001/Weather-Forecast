package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

type WeatherData struct {
	City string  `json:"city"`
	Temp float64 `json:"temp"`
	Desc string  `json:"desc"` // Weather description
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
		// Get 'city' query parameter from the URL
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "City parameter is required", http.StatusBadRequest)
			return
		}

		// Set apiKey from environment variable
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			http.Error(w, "API_KEY not set in .env file", http.StatusInternalServerError)
			return
		}

		// OpenWeatherMap API URL
		apiURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
		resp, err := http.Get(apiURL)
		if err != nil {
			http.Error(w, "Failed to get weather", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Read the response body from OpenWeatherMap
		body, _ := io.ReadAll(resp.Body)

		// Create a struct to hold the API response
		var apiResp struct {
			Name string `json:"name"`
			Main struct {
				Temp float64 `json:"temp"`
			} `json:"main"`
			Weather []struct {
				Description string `json:"description"`
			} `json:"weather"`
		}

		// Decode JSON response into the apiResp struct
		json.Unmarshal(body, &apiResp)

		// Data of the weather
		data := WeatherData{
			City: apiResp.Name,
			Temp: apiResp.Main.Temp,
			Desc: apiResp.Weather[0].Description,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":80", nil)
}

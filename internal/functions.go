package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

// Main page handler
func mainPageHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

// // Function to get coordinates for a given city using OpenWeatherMap API
// func getCoordinates(city string) (float64, float64, error) {
// 	// Set apiKey from environment variable
// 	apiKey := os.Getenv("API_KEY")

// 	// URL for getting coordinates from OpenWeatherMap
// 	apiURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
// 	resp, err := http.Get(apiURL)
// 	if err != nil {
// 		return 0, 0, fmt.Errorf("failed to get coordinates: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body from OpenWeatherMap
// 	body, _ := io.ReadAll(resp.Body)

// 	// Decode JSON response into the GeoResponse struct
// 	var geoResp GeoResponse
// 	json.Unmarshal(body, &geoResp)

// 	// Return coordinates
// 	return geoResp.Coord.Lat, geoResp.Coord.Lon, nil
// }

// Weather API handler
func weatherHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set apiKey from environment variable

		// Get 'city' query parameter from the URL
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "City parameter is required", http.StatusBadRequest)
			return
		}

		// URL for getting weather data from OpenWeatherMap
		apiURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
		resp, err := http.Get(apiURL)
		if err != nil {
			http.Error(w, "Failed to get weather", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Read the response body from OpenWeatherMap
		body, _ := io.ReadAll(resp.Body)

		// Struct for decoding JSON response
		var apiResp struct {
			City string `json:"name"`
			Main struct {
				Temp    float64 `json:"temp"`
				TempMin float64 `json:"temp_min"`
				TempMax float64 `json:"temp_max"`
			} `json:"main"`
			Weather []struct {
				Description string `json:"description"`
			} `json:"weather"`
			Wind struct {
				Speed float64 `json:"speed"`
			} `json:"wind"`
		}

		// Decode JSON response into the apiResp struct
		json.Unmarshal(body, &apiResp)

		// Decoded data from apiResp
		data := WeatherData{
			City:      apiResp.City,
			Desc:      apiResp.Weather[0].Description,
			Temp:      apiResp.Main.Temp,
			TempMin:   apiResp.Main.TempMin,
			TempMax:   apiResp.Main.TempMax,
			WindSpeed: apiResp.Wind.Speed,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

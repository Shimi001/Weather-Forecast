package main

import (
	"log"
	"os"
)

var apiKey string

// Set apiKey from environment variable
func setApiKey() {
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not set in .env file")
	}
}

// key, err := config.GetApiKey()
// if err != nil {
//     http.Error(w, err.Error(), http.StatusInternalServerError)
//     return
// }

// Struct for holding weather data after decoding JSON response
type WeatherData struct {
	City      string  `json:"city"`
	Desc      string  `json:"desc"` // Description
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	WindSpeed float64 `json:"wind_speed"`
}

// Struct for holding coordinates from OpenWeatherMap API response
type GeoResponse struct {
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
}

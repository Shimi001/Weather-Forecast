package main

import (
	"log"
	"os"
)

// Global variable to hold the API key
var apiKey string

// Set apiKey from environment variable
func setApiKey() {
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not set in .env file")
	}
}

// Struct for holding weather data after decoding JSON response
type WeatherData struct {
	MinTemp      float64 `json:"mintemp"`
	MaxTemp      float64 `json:"maxtemp"`
	AvgTemp      float64 `json:"avgtemp"`
	Condition    string  `json:"condition"` // Condition
	WindSpeed    float64 `json:"wind_speed"`
	ChanceOfRain int8    `json:"chance_of_rain"`
}

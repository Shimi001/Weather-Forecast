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
	Temp       float64 `json:"temp"`
	TempMin    float64 `json:"temp_min"`
	TempMax    float64 `json:"temp_max"`
	TempAvg    float64 `json:"temp_avg"`
	Condition  string  `json:"condition"` // Condition
	WindSpeed  float64 `json:"wind_speed"`
	RainChance int8    `json:"rain_chance"`
}

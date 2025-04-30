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

// Current weather API handler
func weatherHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get 'city' query parameter from the URL
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "City parameter is required", http.StatusBadRequest)
			return
		}

		// URL for getting current weather data
		apiURL := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", apiKey, city)
		resp, err := http.Get(apiURL)
		if err != nil {
			http.Error(w, "Failed to get weather", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, _ := io.ReadAll(resp.Body)

		// Struct for decoding JSON response
		var apiResp struct {
			Current struct {
				Temp      float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				WindSpeed float64 `json:"wind_kph"`
			} `json:"current"`
			Forecast struct {
				Forecastday []struct {
					Day struct {
						TempMin    float64 `json:"mintemp_c"`
						TempMax    float64 `json:"maxtemp_c"`
						RainChance int8    `json:"daily_chance_of_rain"`
					} `json:"day"`
				} `json:"forecastday"`
			} `json:"forecast"`
		}

		// Decode JSON response into the apiResp struct
		json.Unmarshal(body, &apiResp)

		// Decoded data from apiResp
		data := WeatherData{
			Temp:       apiResp.Current.Temp,
			TempMin:    apiResp.Forecast.Forecastday[0].Day.TempMin,
			TempMax:    apiResp.Forecast.Forecastday[0].Day.TempMax,
			Condition:  apiResp.Current.Condition.Text,
			WindSpeed:  apiResp.Current.WindSpeed,
			RainChance: apiResp.Forecast.Forecastday[0].Day.RainChance,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

// Forecast weather API handler
func forecastHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

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
						MinTemp      float64 `json:"mintemp_c"`
						MaxTemp      float64 `json:"maxtemp_c"`
						ChanceOfRain int8    `json:"daily_chance_of_rain"`
					} `json:"day"`
				} `json:"forecastday"`
			} `json:"forecast"`
		}

		// Decode JSON response into the apiResp struct
		json.Unmarshal(body, &apiResp)

		// Decoded data from apiResp
		data := WeatherData{
			Temp:         apiResp.Current.Temp,
			MinTemp:      apiResp.Forecast.Forecastday[0].Day.MinTemp,
			MaxTemp:      apiResp.Forecast.Forecastday[0].Day.MaxTemp,
			Condition:    apiResp.Current.Condition.Text,
			WindSpeed:    apiResp.Current.WindSpeed,
			ChanceOfRain: apiResp.Forecast.Forecastday[0].Day.ChanceOfRain,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

// Forecast weather API handler
func forecastHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get 'city' query parameter from the URL
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "City parameter is required", http.StatusBadRequest)
			return
		}

		// URL for getting current weather data
		apiURL := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=7&aqi=no&alerts=no", apiKey, city)
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
			Forecast struct {
				Forecastday []struct {
					Day struct {
						MinTemp      float64 `json:"mintemp_c"`
						MaxTemp      float64 `json:"maxtemp_c"`
						AvgTemp      float64 `json:"avgtemp_c"`
						WindSpeed    float64 `json:"maxwind_kph"`
						ChanceOfRain int8    `json:"daily_chance_of_rain"`
						Condition    struct {
							Text string `json:"text"`
						} `json:"condition"`
					} `json:"day"`
				} `json:"forecastday"`
			} `json:"forecast"`
		}

		// Decode JSON response into the apiResp struct
		json.Unmarshal(body, &apiResp)

		// Decoded data from apiResp
		var data []WeatherData
		for _, day := range apiResp.Forecast.Forecastday {
			data = append(data, WeatherData{
				MinTemp:      day.Day.MinTemp,
				MaxTemp:      day.Day.MaxTemp,
				AvgTemp:      day.Day.AvgTemp,
				Condition:    day.Day.Condition.Text,
				WindSpeed:    day.Day.WindSpeed,
				ChanceOfRain: day.Day.ChanceOfRain,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

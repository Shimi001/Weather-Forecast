import * as DOM from './domElements.js';
import * as func from './functions.js';
import { cityName, currentIndex, forecast } from './config.js';

fetch(`/forecast?city=${cityName}`)
    .then(response => response.json())
    .then(data => {
        console.log(`Forecast data for ${cityName}:`, data);

        forecast.push(...data);
        func.renderForecast(currentIndex, forecast);
    });

// func.buttonEventListeners(); // Add event listeners to the buttons

/*
// Get the current date and day of the week
const today = new Date();
const dayOfWeek = today.getDay(); // Get the current day of the week (0-6)

// Call the functions to set up the page
func.getWeatherInCircle(cityName); // Fetch and display weather data
func.setDayColors(dayOfWeek); // Set the colors for the weather circle and day indicators
func.setDayNames(today, dayOfWeek); // Set the day names and date
*/

/*
import * as func from './functions.js';
import { cityName } from './utils.js';

// Get the current date and day of the week
const today = new Date();
const dayOfWeek = today.getDay(); // Get the current day of the week (0-6)

// Call the functions to set up the page
func.getWeatherInCircle(cityName); // Fetch and display weather data
func.setDayColors(dayOfWeek); // Set the colors for the weather circle and day indicators
func.setDayNames(today, dayOfWeek); // Set the day names and date
*/

fetch('/weather?city=Lviv')
    .then(response => response.json())
    .then(data => {
    console.log("condition " + data.condition);
    console.log("temp " + data.temp); 
    console.log("temp_min " + data.temp_min);
    console.log("temp_max " + data.temp_max);
    console.log("wind_speed " + data.wind_speed)
    console.log("rain_chance " + data.rain_chance);
})

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

fetch('/weather?city=London')
    .then(response => response.json())
    .then(data => {
    console.log(data.city); 
    console.log(data.desc);
    console.log(data.temp); 
    console.log(data.temp_min);
    console.log(data.temp_max);
    console.log(data.wind_speed)
})

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
    console.log("Current weather data for Lviv:");
    console.log("condition " + data.condition);
    console.log("temp " + data.temp); 
    console.log("mintemp " + data.mintemp);
    console.log("maxtemp " + data.maxtemp);
    console.log("wind_speed " + data.wind_speed)
    console.log("chance_of_rain " + data.chance_of_rain);
})

fetch('/forecast?city=Lviv')
    .then(response => response.json())
    .then(data => {
    console.log("Forecast data for Lviv:");
    data.forEach((data, index) => {
        console.log(`Day ${index + 1}:`);
        console.log("condition " + data.condition);
        console.log("avgtemp " + data.avgtemp); 
        console.log("mintemp " + data.mintemp);
        console.log("maxtemp " + data.maxtemp);
        console.log("wind_speed " + data.wind_speed)
        console.log("chance_of_rain " + data.chance_of_rain);
    })
})

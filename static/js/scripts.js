import { cityName, currentIndex, dayOfWeek, forecast } from './config.js';
import { cityElement } from './domElements.js';
import * as func from './functions.js';

export function mainDisplayFuction(currentIndex, dayOfWeek, forecast){
    func.circleColor(dayOfWeek);      // Gives color to the circle
    func.dayIndicators(currentIndex); // Gives color to the indicators
    func.weatherCircle(currentIndex, dayOfWeek, forecast); // Display weather-circle
    func.weatherInfo(currentIndex, forecast);   // Display weather-info
}

fetch(`/forecast?city=${cityName}`)
.then(response => response.json())
.then(data => {
    console.log(`Forecast data for ${cityName}:`, data);
    forecast.push(...data); 
    cityElement.textContent = cityName;      // Display city name
    mainDisplayFuction(currentIndex, dayOfWeek, forecast); // Main display function

    func.arrowsButtons(currentIndex, dayOfWeek, forecast);
});

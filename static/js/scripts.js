import { cityName, dayOfWeek, currentIndex, forecast } from './config.js';
import { cityElement } from './domElements.js';
import * as func from './functions.js';

fetch(`/forecast?city=${cityName}`)
.then(response => response.json())
.then(data => {
    console.log(`Forecast data for ${cityName}:`, data);
    forecast.push(...data); 
    cityElement.textContent = cityName; // Display city name
    func.circleColor(dayOfWeek); // Gives color to the circle
    func.dayIndicators();        // Gives color to the indicators
    func.weatherCircle(forecast, currentIndex, dayOfWeek); // Display weather-circle
    func.weatherInfo(currentIndex, dayOfWeek, forecast);   // Display weather-info

    func.arrowsButtons(currentIndex, dayOfWeek, forecast);
});

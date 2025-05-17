import { today, date, dayColors, japaneseDays } from './config.js';
import * as DOM from './domElements.js';

// Circle Color
export function circleColor(dayOfWeek){
    const circleColor = dayColors[dayOfWeek];
    DOM.circleElement.style.backgroundColor = circleColor;
    DOM.circleElement.style.boxShadow = `0 0 30px ${circleColor}4D`;
}

// Day Indecators
export function dayIndicators(){
    // Day idicators colors
    DOM.dayIndicatorsElement.forEach((indicator, index) => {
        const colorIndex = (today + index) % 7;
        indicator.style.backgroundColor = dayColors[colorIndex];
    });
}

// weather-circle
export function weatherCircle(forecast, currentIndex, dayOfWeek){
    const data = forecast[currentIndex];
    DOM.tempElement.textContent = `${Math.round(data.avgtemp)}°`; // Temperature
    DOM.japaneseDayElement.textContent = japaneseDays[dayOfWeek];
}
 
// weather-info
export function weatherInfo(currentIndex, dayOfWeek, forecast){
    const data = forecast[currentIndex];

    // h1 - day
    DOM.dayElement.textContent = date.toLocaleString('en-US', { weekday: 'long' });

    // h2 - date
    function getOrdinalSuffix(n) {
        if (n > 3 && n < 21) return 'th';
        switch (n % 10) {
          case 1:  return 'st';
          case 2:  return 'nd';
          case 3:  return 'rd';
          default: return 'th';
        }
    }

    const day = date.getDate();
    const month = date.toLocaleString('en-US', { month: 'long' });
    const dayWithSuffix = `${day}${getOrdinalSuffix(day)}`;
    const formatted = `${dayWithSuffix} ${month}`;

    DOM.dateElement.textContent = formatted;

    // weather-details
    DOM.forecastDetailsElement.innerHTML = 
    `${data.condition}. Temperature range from ${Math.round(data.mintemp)}°C to ${Math.round(data.maxtemp)}°C.<br>` +
    `Maximum wind speed ${Math.round(data.wind_speed)} km/h. ${data.chance_of_rain}% daily chance of rain.`;
}

// Prev and next arrows buttons
export const arrowsButtons = (currentIndex, dayOfWeek, forecast) => {
    // nav-arrow next
    DOM.nextButtonElement.addEventListener("click", () => {
        if (currentIndex < forecast.length - 1) {
            currentIndex++;
            date.setDate(date.getDate() + 1);
            dayOfWeek = (dayOfWeek + 1) % 7;
            circleColor(dayOfWeek);
            weatherCircle(forecast, currentIndex, dayOfWeek);
            weatherInfo(currentIndex, dayOfWeek, forecast);
        }
    });

    // nav-arrow prev
    DOM.prevButtonElement.addEventListener("click", () => {
        if (currentIndex > 0) {
            currentIndex--;
            date.setDate(date.getDate() - 1);
            dayOfWeek = (dayOfWeek - 1 + 7) % 7;
            circleColor(dayOfWeek);
            weatherCircle(forecast, currentIndex, dayOfWeek);
            weatherInfo(currentIndex, dayOfWeek, forecast);
        }
    });
}

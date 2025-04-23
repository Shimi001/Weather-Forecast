// const cityElement = document.querySelector('.city-name');

function getWeather(city) {
    fetch(`/weather?city=${city}`)
    .then(response => response.json())
    .then(data => {
        document.getElementById("city").innerText = data.city;
        document.getElementById("temp").innerText = data.temp;
        document.getElementById("desc").innerText = data.desc;
    })
    .catch(error => console.error("Error fetching weather data:", error));
}

const cityName = "Kropyvnytskyi";
// cityElement.textContent = cityName;

getWeather(cityName);

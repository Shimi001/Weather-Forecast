function getWeather(city) {
    fetch(`/weather?city=${city}`)
    .then(response => response.json())
    .then(data => {
        document.getElementById("city").innerText = data.city;
    })
    .catch(error => console.error("Error fetching weather data:", error));
}

getWeather("London");

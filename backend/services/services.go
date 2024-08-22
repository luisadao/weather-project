// services/weather_service.go
package services

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"encoding/json"

	"github.com/luisadao/weather-project/database"
	"github.com/luisadao/weather-project/models"
	"gorm.io/gorm"
)

var cityCodes = map[string]string{
	"Coimbra": "2740636",
	"Faro":    "2268337",
	"Leiria":  "2267094",
	"Lisboa":  "2267056",
	"Porto":   "2735941",
}

func GetCityWeather(cityName string) (*models.WeatherData, error) {
	var weather models.WeatherData

	if err := database.DB.Where("city = ?", cityName).First(&weather).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no such city %v", cityName)
		}
		return nil, fmt.Errorf("error finding city %v: %v", cityName, err)
	}

	if time.Since(weather.Timestamp) > 30*time.Minute {
		weatherFromAPI, err := fetchWeatherFromAPI(cityName)
		if err != nil {
			return nil, err
		}

		if err := database.DB.Model(&models.WeatherData{}).Where("city = ?", weather.City).Updates(weatherFromAPI).Error; err != nil {
			return &weather, fmt.Errorf("error updating city weather: %v", err)
		}
		return weatherFromAPI, nil
	}

	return &weather, nil
}

func GetAllCitiesWeather() ([]models.WeatherData, error) {
	cities := []string{"Lisboa", "Porto", "Coimbra", "Faro", "Leiria"}
	var allWeatherData []models.WeatherData

	for _, city := range cities {
		weather, err := GetCityWeather(city)
		if err == nil {
			allWeatherData = append(allWeatherData, *weather)
		} else {
			fmt.Println("Error fetching weather for city:", city, err)
		}
	}

	return allWeatherData, nil
}

func fetchWeatherFromAPI(cityName string) (*models.WeatherData, error) {
	var weatherFromAPI models.APIResponse
	var myClient = &http.Client{Timeout: 10 * time.Second}
	var weather models.WeatherData

	requestURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?id=%v&APPID=%v", cityCodes[cityName], os.Getenv("API_KEY"))
	resp, err := myClient.Get(requestURL)

	if err != nil {
		return nil, fmt.Errorf("error making http request: %v", err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&weatherFromAPI)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	if len(weatherFromAPI.List) > 0 {
		weather.City = cityName
		weather.Temperature = float32(weatherFromAPI.List[0].Main.Temp - 273.15)
		weather.Description = weatherFromAPI.List[0].Weather[0].Description
		weather.Icon = weatherFromAPI.List[0].Weather[0].Icon
		weather.Timestamp = time.Now()
	} else {
		return nil, fmt.Errorf("no weather data found in API response")
	}
	return &weather, nil
}

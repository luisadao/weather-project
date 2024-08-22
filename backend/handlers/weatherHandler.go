package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisadao/weather-project/controllers"
	"github.com/luisadao/weather-project/models"
)

// GetCityWeather is the route handler to get weather data for a specific city.
func GetCityWeather(c *gin.Context) {
	cityName := c.Param("cityName")

	weather, err := controllers.GetCityWeather(cityName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := models.WeatherDTO{
		City:        weather.City,
		Temperature: weather.Temperature,
		Description: weather.Description,
		Icon:        weather.Icon,
		Timestamp:   weather.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func GetAllCitiesWeatherHandler(c *gin.Context) {
	// Retrieve weather data for all cities
	allWeatherData, err := controllers.GetAllCitiesWeather()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve weather data"})
		return
	}

	var allWeatherDTOs []models.WeatherDTO
	for _, weather := range allWeatherData {
		response := models.WeatherDTO{
			City:        weather.City,
			Temperature: weather.Temperature,
			Description: weather.Description,
			Icon:        weather.Icon,
			Timestamp:   weather.UpdatedAt,
		}
		allWeatherDTOs = append(allWeatherDTOs, response)
	}

	c.JSON(http.StatusOK, allWeatherDTOs)
}

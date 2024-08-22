package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisadao/weather-project/controllers"
	"github.com/luisadao/weather-project/models"
)

func GetCityWeather(c *gin.Context) {
	cityName := c.Param("cityName")

	weather, err := controllers.GetCityWeather(cityName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := models.WeatherResponse{
		City:        weather.City,
		Temperature: weather.Temperature,
		Description: weather.Description,
		Icon:        weather.Icon,
		Timestamp:   weather.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func GetAllCitiesWeatherHandler(c *gin.Context) {
	allWeatherData, err := controllers.GetAllCitiesWeather()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve weather data"})
		return
	}

	var allWeatherResponses []models.WeatherResponse
	for _, weather := range allWeatherData {
		response := models.WeatherResponse{
			City:        weather.City,
			Temperature: weather.Temperature,
			Description: weather.Description,
			Icon:        weather.Icon,
			Timestamp:   weather.UpdatedAt,
		}
		allWeatherResponses = append(allWeatherResponses, response)
	}

	c.JSON(http.StatusOK, allWeatherResponses)
}

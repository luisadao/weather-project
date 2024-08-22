package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/luisadao/weather-project/services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("weather/:cityName", func(c *gin.Context) {
		city := c.Param("cityName")
		weather, err := services.GetCityWeather(city)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, weather)
	})

	r.GET("/weather", func(c *gin.Context) {
		weatherData, err := services.GetAllCitiesWeather()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching weather data"})
			return
		}
		c.JSON(http.StatusOK, weatherData)
	})

	return r
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luisadao/weather-project/database"
	"github.com/luisadao/weather-project/models"
	"gorm.io/gorm"
)

var cityCodes = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Get Cities weather value
	r.GET("weather/:cityName", func(c *gin.Context) {
		city := c.Param("cityName")
		if _, ok := cityCodes[city]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city name"})
			return
		}
		// A weather to hold data from the returned row.
		weather, err := getCityByName(city)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"city": city, "status": "no value"})
		} else {
			c.JSON(http.StatusOK, gin.H{"city": city, "temperature": weather.Temperature, "description": weather.Description, "timestamp": weather.Timestamp, "icon": weather.Icon})
		}
	})

	// Route to get weather for all cities
	r.GET("/weather", func(c *gin.Context) {
		cities := []string{"Lisboa", "Porto", "Coimbra", "Faro", "Leiria"}
		var allWeatherData []models.WeatherData

		for _, city := range cities {
			weather, err := getCityByName(city)
			if err == nil {
				allWeatherData = append(allWeatherData, *weather)
			} else {
				fmt.Println("Error fetching weather for city:", city, err)
			}
		}

		c.JSON(http.StatusOK, allWeatherData)
	})

	return r
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	database.ConnectToDB()
	//Map cityCodes to City
	cityCodes = map[string]string{
		"Coimbra": "2740636",
		"Faro":    "2268337",
		"Leiria":  "2267094",
		"Lisboa":  "2267056",
		"Porto":   "2735941",
	}

	fmt.Println("Connected!")

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run()
}

func getCityByName(cityName string) (*models.WeatherData, error) {
	// Variable to hold the weatherData from the returned row
	var weather models.WeatherData

	if err := database.DB.Where("city = ?", cityName).First(&weather).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no such city %v", cityName)
		}
		return nil, fmt.Errorf("error finding city %v: %v", cityName, err)
	}

	// If the values in the db are older than 30 minutes make new request to OpenWeather
	if time.Since(weather.Timestamp) > 30*time.Minute {
		fmt.Println("Getting new temp from OpenWeather")
		weatherFromAPI, _ := getWeatherFromAPI(cityName)

		// Update the database with the new weather data
		if err := database.DB.Model(&models.WeatherData{}).Where("city = ?", weather.City).Updates(weatherFromAPI).Error; err != nil {
			return &weather, fmt.Errorf("error updating city weather: %v", err)
		}

		return &weather, nil
	}

	return &weather, nil

}

func getWeatherFromAPI(cityName string) (*models.WeatherData, error) {
	var weatherFromAPI models.APIResponse
	var myClient = &http.Client{Timeout: 10 * time.Second}
	var weather models.WeatherData

	requestURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?id=%v&APPID=%v", cityCodes[cityName], os.Getenv("API_KEY"))
	resp, err := myClient.Get(requestURL)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	err = json.NewDecoder(resp.Body).Decode(&weatherFromAPI)
	if err != nil {
		return &weather, fmt.Errorf("error decoding JSON response: %v", err)
	}

	if len(weatherFromAPI.List) > 0 {
		weather.City = cityName
		weather.Temperature = float32(weatherFromAPI.List[0].Main.Temp - 273.15)
		weather.Description = weatherFromAPI.List[0].Weather[0].Description
		weather.Icon = weatherFromAPI.List[0].Weather[0].Icon
		weather.Timestamp = time.Now()
	} else {
		return &weather, fmt.Errorf("no weather data found in API response")
	}
	return &weather, nil
}

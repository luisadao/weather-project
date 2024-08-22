package controllers

import (
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/luisadao/weather-project/database"
	"github.com/luisadao/weather-project/models"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func TestMain(m *testing.M) {
	//Maybe change to MockDB
	database.ConnectToDB()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.WeatherData{})

	m.Run()
}

func TestGetCityWeather(t *testing.T) {
	weather, err := GetCityWeather("Lisboa")

	assert.Nil(t, err)
	fmt.Println(weather)
	assert.Equal(t, "Lisboa", weather.City)
}

func TestGetAllCitiesWeather(t *testing.T) {
	weatherData, err := GetAllCitiesWeather()

	assert.Nil(t, err)
	assert.Equal(t, 5, len(weatherData)) // Expecting 5 cities
}

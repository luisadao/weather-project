package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/luisadao/weather-project/database"
	"github.com/luisadao/weather-project/models"
	"github.com/luisadao/weather-project/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	database.ConnectToDB()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.WeatherData{})

	fmt.Println("Connected!")

	r := routes.SetupRouter()
	r.Run()
}

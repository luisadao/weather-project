package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var cityCodes = make(map[string]string)

type WeatherData struct {
	ID          int64
	City        string
	Temperature float32
	Description string
	Timestamp   time.Time
}

type APIResponse struct {
	List []struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"list"`
	City struct {
		Id int32 `json:"id"`
	} `json:"city"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Get Cities weather value
	r.GET("/:cityName", func(c *gin.Context) {
		city := c.Param("cityName")
		if _, ok := cityCodes[city]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city name"})
			return
		}
		// A weather to hold data from the returned row.
		weather, _ := getCityByName(city)

		c.JSON(http.StatusOK, gin.H{"city": city, "temperature": weather.Temperature})
		/*value, ok := db[city]
		if ok {
			c.JSON(http.StatusOK, gin.H{"city": city, "temperature": weather.Temperature})
		} else {
			c.JSON(http.StatusOK, gin.H{"city": city, "status": "no value"})
		}*/
	})

	return r
}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Capture connection properties.
	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "weather",
		ParseTime: true,
	}
	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	//Map cityCodes to City
	cityCodes = map[string]string{
		"Coimbra": "2740636",
		"Faro":    "2268337",
		"Leira":   "2267094",
		"Lisboa":  "2267056",
		"Porto":   "2735941",
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func getCityByName(cityName string) (WeatherData, error) {
	// Variable to hold the weatherData from the returned row
	var weather WeatherData

	row := db.QueryRow("SELECT city,temperature, description, timestamp FROM weather_data WHERE city = ?", cityName)
	if err := row.Scan(&weather.City, &weather.Temperature, &weather.Description, &weather.Timestamp); err != nil {
		if err == sql.ErrNoRows {
			return weather, fmt.Errorf("no such city %v", cityName)
		}
		return weather, fmt.Errorf("there was an error finding the city %v: %v", cityName, err)
	}

	// If the values in the db are older than 30 minutes make new request to OpenWeather
	if time.Since(weather.Timestamp) > 30*time.Minute {
		fmt.Println("Getting new temp from OpenWeather")

		var weatherFromAPI APIResponse
		var myClient = &http.Client{Timeout: 10 * time.Second}

		requestURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?id=%v&APPID=%v", cityCodes[cityName], os.Getenv("APIKEY"))
		resp, err := myClient.Get(requestURL)

		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			os.Exit(1)
		}

		err = json.NewDecoder(resp.Body).Decode(&weatherFromAPI)
		if err != nil {
			return weather, fmt.Errorf("error decoding JSON response: %v", err)
		}

		fmt.Printf("client: status code: %d", resp.StatusCode)

		if len(weatherFromAPI.List) > 0 {
			weather.City = cityName
			weather.Temperature = float32(weatherFromAPI.List[0].Main.Temp - 273.15)
			weather.Description = weatherFromAPI.List[0].Weather[0].Description
			weather.Timestamp = time.Now()

			// Store the data in the database
			_, err = updateTempCity(weather)
			if err != nil {
				return weather, fmt.Errorf("error updating cache: %v", err)
			}
		} else {
			return weather, fmt.Errorf("no weather data found in API response")
		}

		return weather, err
	}

	return weather, nil

}

func updateTempCity(weather WeatherData) (float32, error) {
	fmt.Printf("\nStoring the value in the database")
	result, err := db.Exec("UPDATE weather_data SET temperature = ? , description = ?, timestamp = ? WHERE city = ?", weather.Temperature, weather.Description, time.Now(), weather.City)
	if err != nil {
		return -1, fmt.Errorf("updateTempCity: %v", err)
	}
	num, err := result.RowsAffected()
	if err != nil || num == 0 {
		return -1, fmt.Errorf("updateTempCity: %v\nRows affected: %d", err, num)
	}
	return weather.Temperature, nil
}

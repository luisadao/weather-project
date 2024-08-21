package models

import "time"

type WeatherData struct {
	ID          int64
	City        string    `json:"city"`
	Temperature float32   `json:"temperature"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Timestamp   time.Time `json:"timestamp"`
}

type APIResponse struct {
	List []struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Icon        string `json:"icon"`
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"list"`
	City struct {
		Id int32 `json:"id"`
	} `json:"city"`
}

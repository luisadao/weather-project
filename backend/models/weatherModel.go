package models

import (
	"time"

	"gorm.io/gorm"
)

type WeatherData struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	City        string    `json:"city"`
	Temperature float32   `json:"temperature"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	UpdatedAt   time.Time `json:"updateAt"`
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

type WeatherDTO struct {
	City        string    `json:"city"`
	Temperature float32   `json:"temperature"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Timestamp   time.Time `json:"timestamp"`
}

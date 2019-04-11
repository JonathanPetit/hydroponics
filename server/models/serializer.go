package models

import (
	"time"
)


type TemperatureResponse struct {
	ID        uint   	`json:"id"`
	Value     float32  	`json:"value"`
	SavedAt   time.Time `json:"savedAt"`
}

type HumidityResponse struct {
	ID        uint   	`json:"id"`
	Value     float32  	`json:"value"`
	SavedAt   time.Time `json:"savedAt"`
}

func (temp *Temperature) Response() TemperatureResponse {
	temperatureResponse := TemperatureResponse{
		ID:        temp.ID,
		Value:     temp.Value,
		SavedAt:   temp.SavedAt,
	}
	return temperatureResponse
}

func (humidity *Humidity) Response() HumidityResponse {
	humidityResponse := HumidityResponse{
		ID:        humidity.ID,
		Value:     humidity.Value,
		SavedAt:   humidity.SavedAt,
	}
	return humidityResponse
}
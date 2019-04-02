package models

import (
	"time"
)


type TemperatureResponse struct {
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
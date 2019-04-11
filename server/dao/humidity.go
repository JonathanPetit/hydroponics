package dao

import (
	"database/sql"
	"fmt"
	"time"
	"net/http"

	"hydroponics/server/models"

	"github.com/gin-gonic/gin"
)

var CREATE_HUMIDITY = `
	INSERT INTO humidity(value, savedAt)
	VALUES($1, now())
	RETURNING id, value, savedAt
`

var GET_HUMIDITIES = `
	SELECT humidity.id, humidity.value, humidity.savedAt FROM humidity
`

func CreateHumidity(c *gin.Context) {
	var humidity models.Humidity
	var id uint
	var value float32
	var savedAt time.Time

	db := c.MustGet("DB").(*sql.DB)
	c.BindJSON(&humidity)

	err := db.QueryRow(CREATE_HUMIDITY, humidity.Value).Scan(&id, &value, &savedAt)
	humidity.SavedAt = savedAt
	humidity.ID = id 
	
	if err != nil {
		fmt.Printf("CREATE_HUMIDITY query error: %v", err)
	}

	c.JSON(http.StatusOK, humidity.Response())
}

func GetHumidities(c *gin.Context) {
	var humidityResponse []models.HumidityResponse

	db := c.MustGet("DB").(*sql.DB)

	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("GET_HUMIDITIES transaction error: %v", err)
	}

	rows, err := tx.Query(GET_HUMIDITIES)
	if err != nil {
		tx.Rollback()
		fmt.Printf("GET_HUMIDITIES query error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uint
		var value float32
		var savedAt time.Time

		if err := rows.Scan(&id, &value, &savedAt); err != nil {
			tx.Rollback()
			fmt.Printf("GET_HUMIDITIES row error: %v", err)
		}

		humidity := models.Humidity{ID: id, Value: value, SavedAt: savedAt}
		humidityResponse = append(humidityResponse, humidity.Response())
	}
	c.JSON(http.StatusOK, humidityResponse)
}


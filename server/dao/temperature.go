package dao

import (
	"database/sql"
	"fmt"
	"time"
	"net/http"

	"hydroponics/server/models"

	"github.com/gin-gonic/gin"
)

var CREATE_TEMPERATURE = `
	INSERT INTO temperature(value, savedAt)
	VALUES($1, now())
	RETURNING id, value, savedAt
`

var GET_TEMPERATURES = `
	SELECT temperature.id, temperature.value, temperature.savedAt FROM temperature
`

func CreateTemperature(c *gin.Context) {
	var temp models.Temperature
	var id uint
	var value float32
	var savedAt time.Time

	db := c.MustGet("DB").(*sql.DB)
	c.BindJSON(&temp)

	err := db.QueryRow(CREATE_TEMPERATURE, temp.Value).Scan(&id, &value, &savedAt)
	temp.SavedAt = savedAt
	temp.ID = id 
	
	if err != nil {
		fmt.Printf("CREATE_TEMPERATURE query error: %v", err)
	}

	c.JSON(http.StatusOK, temp.Response())
}

func GetTemperatures(c *gin.Context) {
	var tempResponse []models.TemperatureResponse

	db := c.MustGet("DB").(*sql.DB)

	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("GET_TEMPERATURES transaction error: %v", err)
	}

	rows, err := tx.Query(GET_TEMPERATURES)
	if err != nil {
		tx.Rollback()
		fmt.Printf("GET_TEMPERATURES query error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uint
		var value float32
		var savedAt time.Time

		if err := rows.Scan(&id, &value, &savedAt); err != nil {
			tx.Rollback()
			fmt.Printf("GET_TEMPERATURES row error: %v", err)
		}

		temp := models.Temperature{ID: id, Value: value, SavedAt: savedAt}
		tempResponse = append(tempResponse, temp.Response())
	}
	c.JSON(http.StatusOK, tempResponse)
}


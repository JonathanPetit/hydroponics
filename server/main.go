package main

import (
	"hydroponics/server/dao"
	"hydroponics/server/database"
	"hydroponics/server/middelware"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	config := database.GetDBConfigFromEnv()

	db := database.ConnectDatabase(config.Name, config.Host, config.User, config.Password, config.Port)

	r := gin.Default()
	r.Use(middelware.Database(db))
	r.Use(cors.Default())
	
	r.POST("/temperature", dao.CreateTemperature)
	r.GET("/temperatures", dao.GetTemperatures)

	r.POST("/humidity", dao.CreateHumidity)
	r.GET("/humidities", dao.GetHumidities)
	//r.GET("/temperatures/:number", dao.GetXTemperatures)
	//r.DELETE("/temperature", dao.DeleteTemperature)
	

	r.Run()
}

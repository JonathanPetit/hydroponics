package main

import (
	"hydroponics/server/dao"
	"hydroponics/server/database"
	"hydroponics/server/middelware"

	"github.com/gin-gonic/gin"
)

func main() {

	config := database.GetDBConfigFromEnv()

	db := database.ConnectDatabase(config.Name, config.Host, config.User, config.Password, config.Port)

	r := gin.Default()
	r.Use(middelware.Database(db))
	
	r.POST("/temperature", dao.CreateTemperature)
	r.GET("/temperatures", dao.GetTemperatures)
	//r.GET("/temperatures/:number", dao.GetXTemperatures)
	//r.DELETE("/temperature", dao.DeleteTemperature)

	r.Run()
}

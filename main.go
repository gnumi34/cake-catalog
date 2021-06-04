package main

import (
	"cake-catalog/database"
	"cake-catalog/models"
	_ "cake-catalog/routers"

	"github.com/astaxie/beego"
)

func main() {
	// Run DB Connection
	database.Connect()
	dbConnector, err := database.Conn.DB()
	if err != nil {
		panic("failed to assign db!")
	}
	defer dbConnector.Close()

	// Run Auto Migration
	database.Conn.AutoMigrate(&models.User{}, &models.Cake{})

	beego.Run()
}

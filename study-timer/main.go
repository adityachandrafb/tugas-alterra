package main

import (
	"study-timer/config"
	"study-timer/route"

	"github.com/labstack/echo/v4"
)

func main() {
	config.initDB()
	config.InitMigration()

	app := echo.New()
	route.NewUser(app)
	app.Start(":8080")
}

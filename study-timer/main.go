package main

import (
	"study-timer/config"
	"study-timer/route"
	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	config.InitMigration()

	app := echo.New()
	route.NewUser(app)
	app.Start(":8080")
}

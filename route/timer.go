package route

import (
	"github.com/labstack/echo"
	"study-timer/controller"
)

func NewRespon(app *echo.Echo) {
	app.GET("/timer", controller.StartTimerController)
	app.GET("/timer", controller.StopTimerController)
	app.GET("/timer", controller.PauseTimerController)
}


package route

import (
	"github.com/labstack/echo"
	"study-timer/controller"
)

func NewTimer(app *echo.Echo) {
	app.GET("/timer", controller.StartTimerController)
	app.GET("/timer", controller.PauseTimeControllerr)
	app.GET("/timer", controller.BreakTimerController)
}
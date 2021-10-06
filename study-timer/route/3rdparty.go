package route

import (
	"study-timer/controller"
	"github.com/labstack/echo"
)

func Audio(app *echo.Echo) {
	app.GET("/music/:audio", controller.GetMusicAudioController)
}
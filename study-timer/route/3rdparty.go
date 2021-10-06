package route

import (
	"study-timer/controller"
	"github.com/labstack/echo"
)

func Audio(e *echo.Echo) {
	e.GET("/music/:audio", controller.GetMusicAudioController)
}
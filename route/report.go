package route

import (
	"study-timer/controller"
	"github.com/labstack/echo"
)

func NewReport(app *echo.Echo) {
	app.GET("/report", controller.GetAllReportController)
	app.POST("/users", controller.CreateReportController)
	app.GET("/report/:id", controller.GetReportByIDController)
	// app.GET("/report/?week", controller.GetReportByWeekController)
}

package route

import (
	"study-timer/controller"
	"github.com/labstack/echo"
)

func NewReport(app *echo.Echo) {
	app.GET("/users", controller.GetAllReportController)
	// app.POST("/users", controller.CreateReportController)
	app.GET("/users/:id", controller.GetReportByIDController)
	app.GET("/users/:week", controller.GetReportByWeekController)
	// app.DELETE("/users/:id", controller.DeleteReportByIDController)
	// app.PUT("/users/:id", controller.UpdateReportByIDController)
}

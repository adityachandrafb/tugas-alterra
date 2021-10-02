package route

import (
	"study-timer/controller"
	"github.com/labstack/echo"
)

func NewGoal(app *echo.Echo) {
	app.GET("/goals", controller.GetAllGoalsController)
	app.POST("/goals", controller.CreateGoalsController)
	app.GET("/goals/:id", controller.GetGoalsByIDController)
	app.DELETE("/goals/:id", controller.DeleteGoalsByIDController)
	app.PUT("/goals/:id", controller.UpdateGoalsByIDController)
}
package route

import (
	"study-timer/controller"
	"study-timer/middleware"

	"github.com/labstack/echo"
)

func NewGoal(app *echo.Echo) {
	// app.GET("/goals", controller.GetAllGoalsController)
	app.POST("/goals", controller.CreateGoalsController)
	app.GET("/goals/:id", controller.GetGoalsByIDController)
	app.DELETE("/goals/:id", controller.DeleteGoalsByIDController, middleware.AuthJWTMiddleware)
	app.PUT("/goals/:id", controller.UpdateGoalsByIDController, middleware.AuthJWTMiddleware)
}
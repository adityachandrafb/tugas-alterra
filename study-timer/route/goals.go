package route

import (
	"study-timer/controller"
	"github.com/labstack/echo"
)

func NewGoal(app *echo.Echo) {
	app.GET("/goals", controller.GetAllUsersController)
	app.POST("/goals", controller.CreateUserController)
	app.GET("/goals/:id", controller.GetUserByIDController)
	app.DELETE("/goals/:id", controller.DeleteUserByIDController)
	app.PUT("/goals/:id", controller.UpdateUserByIDController)
}
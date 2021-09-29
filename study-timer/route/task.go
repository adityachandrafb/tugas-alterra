package route

import (
	"study-timer/controller"
	"github.com/labstack/echo"
)

func NewTask(app *echo.Echo) {
	app.GET("/tasks", controller.GetAllUsersController)
	app.POST("/tasks", controller.CreateUserController)
	app.GET("/tasks/:id", controller.GetUserByIDController)
	app.DELETE("/tasks/:id", controller.DeleteUserByIDController)
	app.PUT("/tasks/:id", controller.UpdateUserByIDController)
}
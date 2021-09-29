package route

import (
	"study-timer/controller"
	"github.com/labstack/echo"
)

func NewUser(app *echo.Echo) {
	app.GET("/users", controller.GetAllUsersController)
	app.POST("/users", controller.CreateUserController)
	app.GET("/users/:id", controller.GetUserByIDController)
	app.DELETE("/users/:id", controller.DeleteUserByIDController)
	app.PUT("/users/:id", controller.UpdateUserByIDController)
}
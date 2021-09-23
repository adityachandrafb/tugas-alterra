package route

import (
	"study-timer/controller"
	"github.com/labstack/echo"
)

func NewUser(app *echo.Echo) {
	app.GET("/users", controller.GetAllUsersController())
	app.POST("/users", controller.CreateUsersController())
	app.GET("/users/:id", controller.GetUsersByIDController)
	app.DELETE("/users/:id", controller.DeleteUsersByIDController)
	app.PUT("/users/:id", controller.UpdateUsersByIDController)
}

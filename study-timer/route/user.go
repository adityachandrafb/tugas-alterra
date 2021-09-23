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

// func NewUser (app *echo.Echo) {

// 	app.GET("/users", controller.GetUsersByIDController)
// 	app.POST("/users", controller.CreateUsersController)
// 	app.GET("/users/:id", controller.GetUsersByIDController)
// 	app.DELETE("/users/:id", controller.DeleteUsersByIDController)
// 	app.PUT("/users/:id", controller.UpdateUsersByIDController)

// }
package route

import (
	
	"study-timer/controller"
	"study-timer/middleware"
	"github.com/labstack/echo"
)

func NewUser(app *echo.Echo) {

	app.GET("/users", controller.GetAllUsersController)
	app.POST("/users", controller.CreateUserController)
	app.POST("/users/login", controller.LoginUserController)
	app.GET("/users/:id", controller.GetUserByIDController)
	app.GET("/users/:id", controller.GetUserDetailController, middleware.JWTAuthMiddleware)
	app.DELETE("/users/:id", controller.DeleteUserByIDController)
	app.PUT("/users/:id", controller.UpdateUserByIDController)
}
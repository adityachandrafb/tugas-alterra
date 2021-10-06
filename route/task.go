package route

import (
	"study-timer/controller"
	"study-timer/middleware"

	"github.com/labstack/echo"
)

func NewTask(app *echo.Echo) {
	app.GET("/tasks", controller.GetAllTasksController, middleware.AuthJWTMiddleware)
	app.POST("/tasks", controller.CreateTaskController)
	app.GET("/tasks/:id", controller.GetTaskByIDController)
	app.DELETE("/tasks/:id", controller.DeleteTaskByIDController)
	app.PUT("/tasks/:id", controller.UpdateTaskByIDController)
}


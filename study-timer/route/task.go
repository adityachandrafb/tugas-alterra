package route

import (
	"study-timer/controller"
	"github.com/labstack/echo"
)

func NewTask(app *echo.Echo) {
	app.GET("/tasks", controller.GetAllTasksController)
	app.POST("/tasks", controller.CreateTaskController)
	app.GET("/tasks/:id", controller.GetTaskByIDController)
	// app.GET("/tasks/:name", controller.GetTaskByNameController)
	app.DELETE("/tasks/:id", controller.DeleteTaskByIDController)
	app.PUT("/tasks/:id", controller.UpdateTaskByIDController)
}


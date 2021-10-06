package controller

import (
	"net/http"
	"study-timer/database"
	"study-timer/model"
	"github.com/labstack/echo"
)

//menampilkan data task
func GetAllTasksController(c echo.Context) error {
	tasks := database.GetAllTasks()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllTasksController",
		"data":    tasks,
	})
}

 //menampilkan data task dengan ID
func GetTaskByIDController(c echo.Context) error {
	id := c.Param("id")
	task := database.GetTaskByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetTaskByIDController",
		"data":    task,
	})
}

 //menghapus data task dg ID
func DeleteTaskByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteTaskByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteTaskByIDController",
	})
}

// mengupdate data task dg ID
func UpdateTaskByIDController(c echo.Context) error {
	id := c.Param("id")

	var task model.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "UpdateTaskByIDController",
			"error":   err.Error(),
		})
	}
	database.UpdateTaskByID(id, task)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetTaskByIDController",
		"data":    task,
	})
}

// menambah task baru
func CreateTaskController(c echo.Context) error {
	var newTask model.Task
	if err := c.Bind(&newTask); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateTaskController",
			"error":   err.Error(),
		})
	}

	newTask = database.CreateTask(newTask)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateTaskController",
		"data":    newTask,
	})
}

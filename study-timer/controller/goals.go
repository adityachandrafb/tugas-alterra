package controller

import (
	"net/http"
	"study-timer/database"
	"study-timer/model"
	"github.com/labstack/echo"
)

//menampilkan semua data goal
func GetAllGoalsController(c echo.Context) error {
	goals := database.GetAllGoals()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllGoalsController",
		"data":    goals,
	})
}

 //menampilkan data goal dengan ID
func GetGoalsByIDController(c echo.Context) error {
	id := c.Param("id")
	goal := database.GetGoalsByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetGoalsByIDController",
		"data":    goal,
	})
}

 //menghapus data goal dg ID
func DeleteGoalsByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteGoalsByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteGoalsByIDController",
	})
}

// mengupdate data goal dg ID
func UpdateGoalsByIDController(c echo.Context) error {
	id := c.Param("id")

	var goal model.Goal
	if err := c.Bind(&goal); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "UpdateGoalsByIDController",
			"error":   err.Error(),
		})
	}
	database.UpdateGoalsByID(id, goal)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetGoalsByIDController",
		"data":    goal,
	})
}

// menambah goal baru
func CreateGoalsController(c echo.Context) error {
	var newGoals model.Goal
	if err := c.Bind(&newGoals); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateGoalsController",
			"error":   err.Error(),
		})
	}

	newGoals = database.CreateGoals(newGoals)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateGoalsController",
		"data":    newGoals,
	})
}



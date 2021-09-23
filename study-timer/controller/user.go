package controller

import (
	"net/http"
	"study-timer/database"
	"study-timer/model"
	"github.com/labstack/echo/v4"
)

//menampilkan data user
func GetAllUsersController(c echo.Context) error {
	users := database.GetUsers
	return c.JSON(http.StatusOK, echo.Map{
		"message": "HalamanIndex",
		"data" : users,
	})	
}

//menampilkan user dengan id
func GetUsersByIDController(c echo.Context) error {
	id := c.Param("id")
	user := database.GetUserByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetUsersByIDController",
		"data" : user,
	})	
}


//delete user
func DeleteUsersByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteUserByID(id)
	return c.JSON(http.StatusOK, echo.Map{
			"message": "DeleteUsersByIDController",
		})
}

//update user
func UpdateUsersByIDControllerr(c echo.Context) error {
	id := c.Param("id")
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "UpdateUserController",
			"error":   err.Error(),
		})
	}

	database.UpdateUserByID(id, user)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "UpdateUserByIDController",
		"data":    user,
	})
	
}


//menambahkan user baru
func CreateUsersController(c echo.Context) error {
	var newUser model.User
		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "CreateUserController",
				"error":   err.Error(),
			})
		}

		newUser = database.CreateUser(newUser)
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"data":    newUser,
		})	
}

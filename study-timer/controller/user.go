package controller

import (
	"net/http"
	"strconv"
	"study-timer/database"
	"study-timer/model"
	"github.com/labstack/echo"
)

//menampilkan data user
func GetAllUsersController(c echo.Context) error {
	users := database.GetUsers()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllUsersController",
		"data":    users,
	})
}

 //menampilkan data user dengan ID
func GetUserByIDController(c echo.Context) error {
	id := c.Param("id")
	user := database.GetUserByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetUserByIDController",
		"data":    user,
	})
}

 //menghapus data user dg ID
func DeleteUserByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteUserByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteUserByIDController",
	})
}

// mengupdate data user dg ID
func UpdateUserByIDController(c echo.Context) error {
	id := c.Param("id")

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "UpdateUserByIDController",
			"error":   err.Error(),
		})
	}
	database.UpdateUserByID(id, user)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetUserByIDController",
		"data":    user,
	})
}

// menambah user baru
func CreateUserController(c echo.Context) error {
	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"error":   err.Error(),
		})
	}

	newUser = database.CreateUser(newUser)
	newUser.Password = ""
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateUserController",
		"data":    newUser,
	})
}

func LoginUserController (c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	token, err := database.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Your Login is Succesfully! Welcome back!",
			"token":   token,
	})
}

func GetUserDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
  
	users, err := database.GetUserDetail(id)
  
	if err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
  
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "status": "login success",
	  "user": users,
	})
  
  }
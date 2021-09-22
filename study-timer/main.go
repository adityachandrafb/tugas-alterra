package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	//"strconv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

type User struct {
	gorm.Model
	Name string //`json:"name" form:"name"`
	Email string //`json:"email" form:"email"`
	Password string //`json:"password" form:"password"`
}

var ( DB *gorm.DB )

func initDB () *gorm.DB {
	dsn := "root:Tya123456@tcp(127.0.0.1:3306)/study-timer?charset=utf8mb4&parseTime=True&loc=Local"
  	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if (err != nil) {
		panic(err.Error())
	}
	return DB
}

func InitMigration (DB *gorm.DB) {
	DB.AutoMigrate(&User{})
}


//menampilkan data user
func GetAllUsersController(DB *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var users []User
		DB.Find(&users)
		
		return c.JSON(http.StatusOK, echo.Map{
			"message": "HalamanIndex",
			"data" : users,

		})
	}	
}

//menampilkan user dengan id
func GetUsersByIDController(DB *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")


		var user User
		DB.Where("id = ?", id).Find(&user)
		if user.ID == 0 {
			return c.JSON(http.StatusNotFound, echo.Map {
				"message" : "Data Tidak Ditemukan",
			} )
		}
		
		return c.JSON(http.StatusOK, echo.Map{
			"message": "GetUsersByIDController",
			"data" : user,

		})
	}	
}


//delete user
func DeleteUsersByIDController(DB *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")

		DB.Where("id = ?", id).Delete(User{})
		
		return c.JSON(http.StatusOK, echo.Map{
			"message": "DeleteUsersByIDController",
		})
	}	
}

//update user
func UpdateUsersByIDControllerr(DB *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")

		DB.Where("id = ?", id).Updates(User{})
		
		return c.JSON(http.StatusOK, echo.Map{
			"message": "UpdateUsersByIDController",
		})
	}	
}


//menambahkan user baru
func CreateUserController(DB *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var newUser User
		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "CreateUserController",
				"error":   err.Error(),
			})
		}
		
		//insert user to mysql
		DB.Create(&newUser)


		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"data":    newUser,
		})
	}	
}

func main() {
	DB := initDB()
	InitMigration(DB)

	app := echo.New()
	app.GET("/users", GetAllUsersController(DB))
	app.POST("/users", CreateUserController(DB))
	app.GET("/users/:id", GetUsersByIDController(DB))
	app.DELETE("/users/:id", DeleteUsersByIDController(DB)) //01.57
	app.PUT("/users/:id", UpdateUsersByIDController(DB))
	// app.POST("/users", HalamanUserTambah)

	// ...
	app.Start(":8080")
}

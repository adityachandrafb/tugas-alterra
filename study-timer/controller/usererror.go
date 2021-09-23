// //menampilkan user dengan id
// func GetUsersByIDController(DB *gorm.DB) func(c echo.Context) error {
// 	return func(c echo.Context) error {
// 		id := c.Param("id")


// 		var user User
// 		DB.Where("id = ?", id).Find(&user)
// 		if user.ID == 0 {
// 			return c.JSON(http.StatusNotFound, echo.Map {
// 				"message" : "Data Tidak Ditemukan",
// 			} )
// 		}
		
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "GetUsersByIDController",
// 			"data" : user,

// 		})
// 	}	
// }


// //delete user
// func DeleteUsersByIDController(DB *gorm.DB) func(c echo.Context) error {
// 	return func(c echo.Context) error {
// 		id := c.Param("id")

// 		DB.Where("id = ?", id).Delete(&User{})
		
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "DeleteUsersByIDController",
// 		})
// 	}	
// }

// //update user
// func UpdateUsersByIDControllerr(DB *gorm.DB) func(c echo.Context) error {
// 	return func(c echo.Context) error {
// 		id := c.Param("id")

// 		DB.Where("id = ?", id).Updates(&User{})
		
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "UpdateUsersByIDController",
// 		})
// 	}	
// }


// //menambahkan user baru
// func CreateUserController(DB *gorm.DB) func(c echo.Context) error {
// 	return func(c echo.Context) error {
// 		var newUser User
// 		if err := c.Bind(&newUser); err != nil {
// 			return c.JSON(http.StatusOK, echo.Map{
// 				"message": "CreateUserController",
// 				"error":   err.Error(),
// 			})
// 		}
		
// 		//insert user to mysql
// 		DB.Create(&newUser)


// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "CreateUserController",
// 			"data":    newUser,
// 		})
// 	}	
// }
 package controller
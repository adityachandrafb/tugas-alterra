package controller

import (
  "study-timer/database"
  "study-timer/model"
  "net/http"
  "github.com/labstack/echo"
)

//menampilkan semua notification
func GetAllNotificationController(c echo.Context) error {
  notifications := database.GetNotifications()
  return c.JSON(http.StatusOK, echo.Map{
    "message": "GetAllNotificationController",
    "data": notifications,
  })
}

// //menampilkan notif by ID
// func GetNotificationByIDController(c echo.Context) error {
//   id := c.Param("id")
//   notification := database.GetNotificationsByID(id)
//   return c.JSON(http.StatusOK, echo.Map{
//     "message": "GetNotificationByIDController",
//     "data": notification.ID,
//   })
// }

//menghapus notif by id
func DeleteNotificationByIDController(c echo.Context) error {
  id := c.Param("id")
  database.DeleteNotificationsByID(id)
  return c.JSON(http.StatusOK, echo.Map{
    "message": "Notification Successfully Deleted",
  })
}

//meng-update notif by id
func UpdateNotificationByIDController(c echo.Context) error {
  id := c.Param("id")

  var notification model.Notification
  if err := c.Bind(&notification); err != nil {
    return c.JSON(http.StatusOK, echo.Map{
      "message": "CreateNotificationController",
      "error": err.Error(),
    })
  }
  database.UpdateNotificationsByID(id, notification)
  return c.JSON(http.StatusOK, echo.Map{
    "message": "Notification Successfully Updated",
  })
}

//membuat notif baru
func CreateNotificationController(c echo.Context) error {
  var newNotification model.Notification
  if err := c.Bind(&newNotification); err != nil {
    return c.JSON(http.StatusOK, echo.Map{
      "message": "CreateNotificationController",
      "error": err.Error(),
    })
  }

  newNotification = database.CreateNotifications(newNotification)
  return c.JSON(http.StatusOK, echo.Map{
    "message": "CreateNotificationController",
    "data": newNotification,
  })
}
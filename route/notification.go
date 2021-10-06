package route

import (
	"study-timer/controller"
	"github.com/labstack/echo"
)

func NewNotification(app *echo.Echo) {
	app.GET("/notif", controller.GetAllNotificationController)
	app.POST("/notif", controller.CreateNotificationController)
	// app.GET("/notif/:id", controller.GetNotificationsByIDController)
	app.DELETE("/notif/:id", controller.DeleteNotificationByIDController)
	app.PUT("/notif/:id", controller.UpdateNotificationByIDController)
}
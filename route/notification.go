package route

import (
	"study-timer/controller"
	"study-timer/middleware"
	"github.com/labstack/echo"
)

func NewNotification(app *echo.Echo) {
	app.GET("/notif", controller.GetAllNotificationController)
	app.POST("/notif", controller.CreateNotificationController)
	// app.GET("/notif/:id", controller.GetNotificationsByIDController)
	app.DELETE("/notif/:id", controller.DeleteNotificationByIDController, middleware.AuthJWTMiddleware)
	app.PUT("/notif/:id", controller.UpdateNotificationByIDController, middleware.AuthJWTMiddleware)
}
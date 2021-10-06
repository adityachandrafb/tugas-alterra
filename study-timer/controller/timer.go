package controller

import (
	"net/http"
	"study-timer/database"
	"github.com/labstack/echo"
)

//menampilkan status timer mulai
 func StartTimerController(c echo.Context) error {
	status := c.Param("start")
	database.StartTimer(status)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Time to study!",
	})
}


//menampilkan status timer stop
func StopTimerController(c echo.Context) error {
	status := c.Param("stop")
	database.StopTimer(status)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Finish! You are doing great job!",
	})
}

//menampilkan status timer pause
func PauseTimerController(c echo.Context) error {
	status := c.Param("pause")
	database.PauseTimer(status)
	return c.JSON(http.StatusOK, echo.Map{
		"message" : "You're not done yet!Lets go back to study!",
	})
}


















//pomodoro timer
// func StartPomodoro(cycles int64) {
// 	//Belajar 25 menit
// 	var minutesPassed int64 = 0
// 	for i := 0; i < 25; i++ {
// 		time.Sleep(time.Minute * 1)
// 		minutesPassed++
// 		t := 25 - minutesPassed
// 		return StartStudy
// 	}
// 	return PauseTimer

// 	// Break 5 menit
// 	minutesPassed = 0
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Minute * 1)
// 		minutesPassed++
// 		t := 5 - minutesPassed
// 		return BreakStudy
// 	}
// 	return StartStudy
// }


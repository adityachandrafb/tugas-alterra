package controller

import ()

)


//pomodoro timer
func StartPomodoro(cycles int64) {
	//Belajar 25 menit
	var minutesPassed int64 = 0
	for i := 0; i < 25; i++ {
		time.Sleep(time.Minute * 1)
		minutesPassed++
		t := 25 - minutesPassed
		return StartStudy
	}
	return PauseTimer

	// Break 5 menit
	minutesPassed = 0
	for i := 0; i < 5; i++ {
		time.Sleep(time.Minute * 1)
		minutesPassed++
		t := 5 - minutesPassed
		return BreakStudy
	}
	return StartStudy
}


//memulai timer 25 menit belajar
func StartTimerController (c echo.Context) error {
	if StartStudy == true {
		return c.JSON(http.StatusOK, echo.Map{
			"message" : "Time to study!",
		})
}


//mengetahui sisa waktu timer
func PauseTimerController(c echo.Context) error {
	remainTime = 25 - const.time.Duration
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Lets go back to study!",
		"Remain timer" : remainTime
	})
}

//memulai timer 5 menit istirahat
func BreakTimerController(c echo.Context) error {
	if BreakStudy == true {
		return c.JSON(http.StatusOK, echo.Map{
			"message" : "Enjoy your break time!",
		})
}

//menyelsaikan timer dan mengirim ke report
func FinishTimer(c echo.Context) error {
	var FinishStudy = 4 * StartPomodoro
	switch FinishStudy {
	case < 4 :
		return c.JSON(http.StatusOK, echo.Map{
			"message": "You are almost done! Keep going", 
		)}
	case 4 : 
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Finish! You are doing great job!",
		)}
	default :
	}
	//save ke report ???
}
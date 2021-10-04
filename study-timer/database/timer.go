package database

import (
	"study-timer/config"
	"study-timer/model"
)

func StartTimer() model.Pomodoro {
	var start model.Pomodoro
	config.DB.Where("min is 1").Find(&start)
	return start
}

func PauseTimer(id string) model.Pomodoro {
	var pause model.Pomodoro
	config.DB.Where("min is not 25", id).Find(&pause)
	return pause
}

func BreakTimer(id string) model.Pomodoro {
	var breaks model.Pomodoro
	config.DB.Where("min is 25", id).Find(&breaks)
	return breaks
}

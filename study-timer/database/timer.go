package database

import (
	"study-timer/config"
	"study-timer/model"
)

func StartTimer(status string) model.Respon {
	var start model.Respon
	config.DB.Find(&start)
	return start
}

func StopTimer(status string) model.Respon {
	var stop model.Respon
	config.DB.Find(&stop)
	return stop
}

func PauseTimer(status string) model.Respon {
	var pause model.Respon
	config.DB.Find(&pause)
	return pause
}

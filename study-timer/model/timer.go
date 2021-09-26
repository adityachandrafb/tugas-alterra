package model

import (
	"time"
	"gorm.io/gorm"
)

const (
	Duration = 25 * time.Minute
	Shortbreak = 5 * time.Minute
	Longbreak = 30 * time.Minute
)

var MyTask []*model.Tasks `gorm:"foreignkey:username;references:ID"` //One to One Tasks dengan refrensi ID, array karena task yg dikerjakan bisa lebih dari satu

type Timer struct {
	gorm.Model
	ID      	uint   `gorm:"foreignkey"`
	Duration     time.Duration
	StartAt      time.Time
	Paused       time.Time
	Resumed      time.Time
	StopAt       time.Time
	Finished     bool
}

// type Timer struct {
// 	gorm.Model
// 	ID           string
// 	Duration     time.Duration
// 	StartAt      time.Time
// 	EndAt        time.Time
// 	StopAt       null.Time
// 	LastPauseAt  null.Time
// 	LastResumeAt null.Time
// 	FinishAt     null.Time
// 	Completed    bool
// }
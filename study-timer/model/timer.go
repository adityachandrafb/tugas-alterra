package model

import (
	"time"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)
type Pomodoro struct {
	Timer time.Time
	Loop  int
	Min   int
	Sec   int
}

const (
	Duration = 25 * time.Minute
	Shortbreak = 5 * time.Minute
	Longbreak = 15 * time.Minute
	LongbreakInterval = 4 
)

type Timer struct {
	gorm.Model
	ID				uint   		`gorm:"foreignkey"`
	Duration     time.Duration	`json:"duration"`
	StartAt      time.Time		`json:"start_at"`
	// PauseAt      time.Now		`json:"pause_at"`
	// ResumeAt     time.Now		`json:"pause_at"`
	StopAt       null.Time		`json:"stop_at"`
	Completed    bool			`json:"completed"`
}


//timer perlu gak sih masukkin json?
package model

import ()
// 	"time"
// 	"gorm.io/gorm"
// )

// const (
// 	Duration = 25 * time.Minute
// 	Shortbreak = 5 * time.Minute
// 	Longbreak = 30 * time.Minute
// )

// var MyTask []*model.Tasks //`gorm:"foreignkey:username;references:ID"` //One to One Tasks dengan refrensi ID, array karena task yg dikerjakan bisa lebih dari satu

// type Timer struct {
// 	gorm.Model
// 	ID				uint   		`gorm:"foreignkey"`
// 	Duration     time.Duration	`json:"duration"`
// 	StartAt      time.Time		`json:"start_at"`
// 	EndAt        time.Time		`json:"end_at"`
// 	StopAt       null.Time		`json:"stop_at"`
// 	LastPauseAt  null.Time		`json:"last_pause_at"`
// 	LastResumeAt null.Time		`json:"last_resume_at"`
// 	FinishAt     null.Time		`json:"finish_at"`
// 	Completed    bool			`json:"completed"`
// }

//timer perlu gak sih masukkin json?
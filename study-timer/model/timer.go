package model

import (
	// "time"
	// "gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Respon struct {
	gorm.Model
	Start string
	Pause string
	Stop string
}



















// const (
// 	Duration = 25 * time.Minute
// 	Shortbreak = 5 * time.Minute
// 	Longbreak = 15 * time.Minute
// 	LongbreakInterval = 4 
// )


// type Timer struct {
// 	gorm.Model
// 	ID				uint   		`gorm:"foreignkey"`
// 	Duration    time.Duration	`json:"duration"`
// 	StartAt     time.Time		`json:"start_at"`
// 	StopAt      null.Time		`json:"stop_at"`
// 	Status    	string			`json:"status"`
// }

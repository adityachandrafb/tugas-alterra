package model

import (
	"time"
	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	ID       		uint   		`gorm:"primarykey"`
	TargetDay		int			`json:"target_day"`
	TargetTime		int			`json:"target_time"`
	StartDate 		time.Time	`json:"start_date"`
	FinishDate		time.Time	`json:"finish_date"`
	UserID			uint		`json:"user_id"`	
	User			*User 		`json:"user,omitempty"`
}
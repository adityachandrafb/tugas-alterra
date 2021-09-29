package model

import (
	"gorm.io/gorm"
)

type Goals struct {
	gorm.Model
	ID       		uint   		`gorm:"primarykey"`
	TargetDay		int			`json:"target_day"`
	TargetTime		int			`json:"target_time"`
	StartDate 		date.date	//kenapa ya?
	FinishDate		null.date
	
}
package model

import "gorm.io/gorm"

type Goals struct {
	gorm.Model
	ID       		uint   `gorm:"primarykey"`
	TargetDay		int
	TargetTime		int
	Start 			date 
	Finish			date
}
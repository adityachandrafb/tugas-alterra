 package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID      	uint   	`gorm:"foreignkey"`
	Name    	string 	`json:"name"`
	Deadline    string	`json:"date"` //date
	Timer 		*Timer 	
	Priority	bool  	`json:"priority"`//penting & mendesak atau tidak
}
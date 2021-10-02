package model

import (
	"time"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID      	uint   	`gorm:"foreignkey"`
	Name    	string 	`json:"name"`
	Deadline    time.Time	`json:"deadline"`
	// Timer 		*Timer 	`json:"timer,omitempty"`
	Priority	bool  	`json:"priority"`//penting & mendesak atau tidak
	UserID		uint	`json:"user_id"`	
	User		*User 	`json:"user,omitempty"`
}

type Tasks []Task
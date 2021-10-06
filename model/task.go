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
	UserID		uint	`json:"user_id"`	
	User		*User 	`json:"user,omitempty"`
}
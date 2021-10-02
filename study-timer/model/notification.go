package model

import (
	"time"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID       		uint   		`gorm:"primarykey"`
	reminder 		string		
	time			time.Time
	UserID			uint		`json:"user_id"`	
	GoalID			uint 		`json:"goal,omitempty"`
	TaskID			uint		`json:"task,omitempty"`
}

//notif many to one ke user, goal, task
// setiap user, goal, task punya banyak notification
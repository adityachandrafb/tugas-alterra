package model

import (
	"gorm.io/gorm"
	"time"
)
type Report struct {
	gorm.Model
	Week			time.Time	`json:"week"`
	Goal 			*Goal		`json:"goals,omitempty"`
	Achivement 		int 		//percentage
	UserID			uint		`json:"user_id"`	
	User			*User 		`json:"user,omitempty"`
}

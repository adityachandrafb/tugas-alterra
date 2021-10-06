package model

import (
	"gorm.io/gorm"
	//"time"
)
type Report struct {
	gorm.Model
	Week			int		`json:"week"`
	Completed 		bool 		`json:"completed,omitempty"`
	Goal			*Goal 		`json:"goal,omitempty"`
	GoalID			uint		`json:"user_id"`
}

package model

import "gorm.io/gorm"

type Tasks struct {
	gorm.Model
	ID      	uint   `gorm:"foreignkey"`
	Name    	string `json:"name"`
	Deadline    string`json:"date"` //date
	Timer 		string `json:"time"`
	Priority	bool  //urgent atau tidak
}
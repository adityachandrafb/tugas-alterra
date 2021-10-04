package model

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      	 uint   `gorm:"primarykey"`
	Username    string 	`json:"username"`
	Email    	string	`json:"email"`
	Password 	string 	`json:"password"`
	Task		[]*Task `json:"task"`
	//CreatedAt time.Time `json:"created_at"`
	DeletedAt	*time.Time `json:"deleted_at"`
	Token 		string 		`json:"token"`
}

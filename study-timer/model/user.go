package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      	 uint   `gorm:"primarykey"`
	Username    string 	`json:"username"`
	Email    	string	`json:"email"`
	Password 	string 	`json:"password"`
}

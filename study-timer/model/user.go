package model

import (
	
)

type User struct {
	//gorm.Model
	ID      	 uint   `gorm:"primarykey"`
	Username    string `json:"name"`
	Email    	string `json:"email"`
	Password 	string `json:"password"`
}

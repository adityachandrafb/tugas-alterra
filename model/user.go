package model

import (
	"time"
	//"gorm.io/gorm"
)


type User struct {
	//gorm.Model
	ID      	 uint   `gorm:"primarykey"`
	Username    string 	`json:"username"`
	Email    	string	`json:"email"`
	Password 	string 	`json:"password"`
	Task		[]*Task  `json:"task"`
	Goal		[]*Goal  `json:"goal"`
	CreatedAt 	time.Time `json:"createdAt"`
	DeletedAt	*time.Time `json:"deletedAt"`
	Token 		string 	    `json:"token,omitempty"`
}

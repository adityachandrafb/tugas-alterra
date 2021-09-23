package database

import (
	"study-timer/config"
	"study-timer/model"
)

func GetUsers() []model.User {
	var users []model.User
	config.DB.Find(&users)
	return users
}

func GetUserByID(id string) model.User {
	var user model.User
	config.DB.Where("id = ?", id).Find(&user)
	return user
}

func CreateUser(user model.User) model.User {
	config.DB.Create(&user)
	return user
}

func DeleteUserByID(id string) {
	var user model.User
	config.DB.Where("id = ?", id).Delete(&user)
}

func UpdateUserByID(id string, user model.User) {
	config.DB.Where("id = ?", id).Updates(&user)
}

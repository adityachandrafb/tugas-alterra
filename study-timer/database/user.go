package database

import (
	"errors"
	"study-timer/config"
	"study-timer/middleware"
	"study-timer/model"
	"time"
)

func GetUsers() []model.User {
	var users []model.User
	config.DB.Preload("Task").Where("deleted_at is null").Find(&users)
	return users
}

func GetUserByID(id string) model.User {
	var user model.User
	config.DB.Where("id = ?", id).Preload("Task").Find(&user)
	return user
}

func CreateUser(user model.User) model.User {
	config.DB.Create(&user)
	return user
}

func DeleteUserByID(id string) {
	var user model.User
	deleteUser := time.Now()
	user.DeletedAt = &deleteUser
	config.DB.Where("id = ?", id).Updates(&user)
}

func UpdateUserByID(id string, user model.User) {
	config.DB.Where("id = ?", id).Updates(&user)
}

func IsValid (nm string, em string) bool {
	var user model.User
	if err := config.DB.Where("email = ?", em).Find(&user).Error; err != nil {
		return false
	}
	return nm == user.Username
}

func GetUserDetail(userId int)(interface{}, error) {
	var user model.User

	if e := config.DB.Find(&user,userId).Error; e != nil {
		return nil, e
	}
	return user,nil
}

func LoginUser (user *model.User) (interface{}, error) {
	var userFound model.User
	var err error
	if err = config.DB.Where("email = ?", user.Email). First(userFound).Error; err != nil {
		return nil, err
	} 

	if user.Password != userFound.Password {
		return nil, errors.New("User not found. Your username or password are incorret")
	}
	user.Token, err = middleware.CreateToken(int(user.ID)) 
	if err != nil {
		return nil, err
	}
	return user.Token, nil
}

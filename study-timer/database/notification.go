package database

import (
	"study-timer/config"
	"study-timer/model"
)

func GetNotifications() []model.Notification {
	var goal []model.Notification
	config.DB.Find(&goal)
	return goal
}

func CreateNotifications(goal model.Notification) model.Notification {
	config.DB.Create(&goal)
	return goal
}

func DeleteNotificationsByID(id string) {
	var goal model.Notification
	config.DB.Where("id = ?", id).Delete(&goal)
}

func UpdateNotificationsByID(id string, goal model.Notification) {
	config.DB.Where("id = ?", id).Updates(&goal)
}
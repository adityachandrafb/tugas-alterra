package database

import (
	"study-timer/config"
	"study-timer/model"
)

func GetAllTasks() []model.Task {
	var task []model.Task
	config.DB.Find(&task)
	return task
}

func GetTaskByID(id string) model.Task {
	var task model.Task
	config.DB.Where("id = ?", id).Find(&task)
	return task
}

func CreateTask(task model.Task) model.Task {
	config.DB.Create(&task)
	return task
}

func DeleteTaskByID(id string) {
	var task model.Task
	config.DB.Where("id = ?", id).Delete(&task)
}

func UpdateTaskByID(id string, task model.Task) {
	config.DB.Where("id = ?", id).Updates(&task)
}

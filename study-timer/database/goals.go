package database

import (
	"study-timer/config"
	"study-timer/model"
)

func GetAllGoals() []model.Goal {
	var goal []model.Goal
	config.DB.Find(&goal)
	return goal
}

func GetGoalsByID(id string) model.Goal {
	var goal model.Goal
	config.DB.Where("id = ?", id).Find(&goal)
	return goal
}

func CreateGoals(goal model.Goal) model.Goal {
	config.DB.Create(&goal)
	return goal
}

func DeleteGoalsByID(id string) {
	var goal model.Goal
	config.DB.Where("id = ?", id).Delete(&goal)
}

func UpdateGoalsByID(id string, goal model.Goal) {
	config.DB.Where("id = ?", id).Updates(&goal)
}

package database

import (
	"study-timer/config"
	"study-timer/model"
)

func GetAllGoals() []model.Goals {
	var goal []model.Goals
	config.DB.Find(&goal)
	return goal
}

func GetGoalsByID(id string) model.Goals {
	var goal model.Goals
	config.DB.Where("id = ?", id).Find(&goal)
	return goal
}

func CreateGoals(goal model.Goals) model.Goals {
	config.DB.Create(&goal)
	return goal
}

func DeleteGoalsByID(id string) {
	var goal model.Goals
	config.DB.Where("id = ?", id).Delete(&goal)
}

func UpdateGoalsByID(id string, goal model.Goals) {
	config.DB.Where("id = ?", id).Updates(&goal)
}

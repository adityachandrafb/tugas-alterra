package model

import (
	"gorm.io/gorm"
	"date" //harus impor package
)
type Report struct {
	gorm.Model
	Week	date.Date
	Goal 	*Goals
	Achivement int //percentage
}

package model

import (
	"gorm.io/gorm"
	//"time"
)
type Report struct {
	gorm.Model
	Week			int			`json:"week"`
}

package model

import (
	"time"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID       		uint   		`gorm:"primarykey"`
	NotifName 		string		`json:"notif_id"`
	time			time.Time
}

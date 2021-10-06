package database

import (
	"study-timer/config"
	"study-timer/model"
)

//cari semua report
func GetAllReport() []model.Report {
	var report []model.Report
	config.DB.Find(&report)
	return report
}
//cari report berdasarkan id
func GetReportByID(id string) model.Report {
	var report model.Report
	config.DB.Preload("Goal").Where("id = ?", id).Find(&report)
	return report
}

//membuat report
func CreateReport(goal string) model.Report {
	var report model.Report
	config.DB.Create(&report)
	return report
}

package database

import (
	"study-timer/config"
	"study-timer/model"
)

//semua report
func GetAllReport() []model.Report {
	var report []model.Report
	config.DB.Find(&report)
	return report
}

//cari report berdasarkan minggu
func GetReportByDate(date string) model.Report {
	var report model.Report
	config.DB.Where("date = ?", date).Find(&report)
	return report
}

//cari report berdasarkan id
func GetReportByID(id string) model.Report {
	var report model.Report
	config.DB.Where("id = ?", id).Find(&report)
	return report
}

//membuat report
func CreateReport(goal string) model.Report {
	var report model.Goals
	config.DB.Create(&report)
	return report
}

// update report setiap ada goals baru tercapai
// func UpdateReportByID(id string, task model.Task) {
// 	config.DB.Where("id = ?", id).Updates(&task)
// }


package controller

import (
	"net/http"
	"study-timer/database"
	// "study-timer/model"
	"github.com/labstack/echo"
)

//menampilkan data report
func GetAllReportController(c echo.Context) error {
	report := database.GetAllReport()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllReportController",
		"data":    report,
	})
}

 //menampilkan data report dengan ID
func GetReportByIDController(c echo.Context) error {
	id := c.Param("id")
	report := database.GetReportByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetReportByIDController",
		"data":    report,
	})
}

 //menampilkan data report dengan week/minggu
 func GetReportByWeekController(c echo.Context) error {
	week := c.Param("week")
	report := database.GetReportByID(week)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetReportByWeekController",
		"data":    report,
	})
}

// // menambah report baru otomatis dari goal atau task
// func CreateReportController(c echo.Context) error {
// 	var newReport model.Report
// 	if err := c.Bind(&newReport); err != nil {
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "CreateReportController",
// 			"error":   err.Error(),
// 		})
// 	}

// 	newReport = database.CreateReport(newReport)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "CreateReportController",
// 		"data":    newReport,
// 	})
// }

// // mengupdate data report dg ID
// func UpdateReportByIDController(c echo.Context) error {
// 	id := c.Param("id")

// 	var report model.Report
// 	if err := c.Bind(&report); err != nil {
// 		return c.JSON(http.StatusOK, echo.Map{
// 			"message": "UpdateReportByIDController",
// 			"error":   err.Error(),
// 		})
// 	}
// 	database.UpdateReportByID(id, Report)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "GetReportByIDController",
// 		"data":    report,
// 	})
// }
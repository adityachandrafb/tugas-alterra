package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"study-timer/model"
)

var ( 
	DB *gorm.DB )

// fungsi connect ke mysql
func InitDB() {
	dsn := "root:Tya123456@tcp(127.0.0.1:3306)/study-timer?charset=utf8mb4&parseTime=True&loc=Local"
  	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if (err != nil) {
		panic(err)
	}
}

// fungsi bikin tabel otomatis
func InitMigration() {
	DB.AutoMigrate(&model.User{})
}
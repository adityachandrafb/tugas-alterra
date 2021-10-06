package config

import (
	"context"
	"study-timer/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBLog *mongo.Client
	

// fungsi connect ke mysql
func InitDB() {
	dsn := "root:Tya123456@tcp(127.0.0.1:3306)/study-timer?charset=utf8mb4&parseTime=True&loc=Local"
  	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if (err != nil) {
		panic(err)
	}
}

//fungsi connect mongodb
func InitLog() {
	var err error
	DBLog, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/study-timer"))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err = DBLog.Connect(ctx)
	if err != nil {
		panic(err)
	}

	DBLog.ListDatabaseNames(ctx, bson.M{})
}


// fungsi bikin tabel otomatis
func InitMigration() {
	DB.AutoMigrate(
		&model.User{},
		&model.Task{},
		&model.Goal{},
		&model.Notification{},
		&model.Respon{},

	)
}

//coba karena masih error

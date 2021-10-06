package middleware

import (
	"study-timer/config"
	//"study-timer/model"
	"fmt"
	//"net/http"
	"time"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	collection := config.DBLog.Database("study-timer").Collection("Logger")

	return func(c echo.Context) error {

		log := bson.M{
			"time":   time.Now(),
			"method": c.Request().Method,
			"path":   c.Path(),
		}

		collection.InsertOne(c.Request().Context(), log)

		fmt.Println(log)

		return next(c)
	}
}


















// func Log(next echo.HandlerFunc) echo.HandlerFunc {
// 	coll := config.DBLog.Database("study-timer").Collection("logs")

// 	return func(c echo.Context) error {
// 		data := new(model.User)
// 		if c.Request().Method != http.MethodGet {
// 			if err := c.Bind(&data); err != nil {
// 				fmt.Println(err)
// 			}
// 		}

// 		log := bson.M{
// 			"time":    time.Now(),
// 			"method":  c.Request().Method,
// 			"path":    c.Path(),
// 			"request": data,
// 		}

// 		response := next(c)
// 		log["response"] = c.Response().Status
// 		coll.InsertOne(c.Request().Context(), log)
// 		fmt.Println(log)
// 		return response
// 	}
// }

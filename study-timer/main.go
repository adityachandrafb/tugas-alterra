package main

import (
	"fmt"
	"net/http"
	"study-timer/config"
	"study-timer/database"
	"study-timer/middleware"
	"study-timer/route"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)
func HandlerTest(c echo.Context) error {
	return c.String(http.StatusOK, "Berhasil")
}

func HandlerLogin(c echo.Context) error {
	user := struct {
		Email    string
		Password string
	}{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	isValid := database.IsValid(user.Email, user.Password)
	if !isValid {
		return c.String(http.StatusBadRequest, "email atau password salah")
	}

	claims := jwt.MapClaims{}
	claims["userId"] = user.Email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("asdasdasd"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, tokenString)
}

func main() {
	config.InitDB()
	config.InitLog()
	config.InitMigration()

	app := echo.New()
	app.Use(middleware.Log)
	
	app.GET("/test", HandlerTest)

	app.GET(
		"/",
		func(c echo.Context) error {
			email := c.Get("email")
			return c.String(http.StatusOK, fmt.Sprintf("selamat datang %s", email))
		},
		middleware.JWTAuthMiddleware,
	)

	app.POST("/login", HandlerLogin)


	route.NewUser(app)
	route.NewTask(app)
	route.NewGoal(app)
	route.NewNotification(app)
	// route.NewTimer(app)
	
	app.Start(":8080")
}

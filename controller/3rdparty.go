package controller

import (
	"study-timer/database"
	"net/http"
	"github.com/labstack/echo"
)

func GetMusicAudioController(c echo.Context) error {
	result := database.GetAudio()

	return c.JSON(http.StatusOK, echo.Map{
		"message":    "Playing music!",
		"dictionary": result,
	})
}

// func main() {
// 	result := GetArtist()
// 	fmt.Println(result.AAAA)
// }
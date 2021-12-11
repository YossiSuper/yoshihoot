package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

//トップページ
func TopPage(c echo.Context) error {
	data := struct {
		CountOfPlayers int
		CountOfRooms   int
		UserID         string
	}{
		CountOfPlayers: 20,
		CountOfRooms:   2,
		UserID:         "4",
	}
	return c.Render(http.StatusOK, "index", data)
}

func CreateGame(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func JoinGame(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func DeleteGame(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func GetGames(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func LogIn(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}

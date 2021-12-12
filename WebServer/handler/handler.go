package handler

import (
	"net/http"

	"yoshihoot_WebServer/api"

	"github.com/labstack/echo"
)

//トップページ
func TopPage(c echo.Context) error {
	data := struct {
		CountOfPlayers int
		CountOfRooms   int
	}{
		CountOfPlayers: 20,
		CountOfRooms:   2,
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

func GetServerStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func LogIn(c echo.Context) error {
	return c.JSON(http.StatusOK, api.LogIn(c.QueryParam("name")))
}

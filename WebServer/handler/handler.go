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

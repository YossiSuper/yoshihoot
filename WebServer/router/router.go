package router

import (
	"yoshihoot_WebServer/handler"

	"github.com/labstack/echo"
)

func SetRouter(e *echo.Echo) {
	e.Static("/public", "public")

	e.GET("/", handler.TopPage)

	api := e.Group("/api")
	api.GET("/LogIn", handler.LogIn)
	api.GET("/CreateGame", handler.CreateGame)
	api.GET("/JoinGame", handler.JoinGame)
	api.GET("/DeleteGame", handler.DeleteGame)
	api.GET("/GetGames", handler.GetGames)
	api.GET("/GetServerStatus", handler.GetServerStatus)
}

package router

import (
	"yoshihoot_WebServer/handler"

	"github.com/labstack/echo"
)

func SetRouter(e *echo.Echo) {
	e.Static("/public", "public")

	e.GET("/", handler.TopPage)
}

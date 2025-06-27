package main

import (
	"bellog/src/components"

	"github.com/labstack/echo/v4"
)

func main() {
	home := components.Home("test")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return home.Render(c.Request().Context(), c.Response().Writer)
	})
	e.Logger.Fatal(e.Start(":9001"))
}

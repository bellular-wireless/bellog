package handlers

import (
	"bellog/src/components"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	home := components.Index(components.Home())
	return Render(c, 200, home)
}

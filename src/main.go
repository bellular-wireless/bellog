package main

import (
	"bellog/src/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.HomeHandler)
	e.GET("/articles", handlers.ArticleListHandler)

	e.Logger.Fatal(e.Start(":9001"))
}

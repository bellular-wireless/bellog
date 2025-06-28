package handlers

import (
	"bellog/src/components"

	"github.com/labstack/echo/v4"
)

func ArticleListHandler(c echo.Context) error {
	article := components.Article{Content: "this is a test article", Title: "Test Article", Author: "John Doe", Date: "2023-10-01"}
	articleList := components.ArticleList([]components.Article{article, article, article, article})

	return Render(c, 200, articleList)
}

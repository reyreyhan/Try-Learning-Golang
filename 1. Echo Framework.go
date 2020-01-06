package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func main() {
	route := echo.New()

	route.GET("/", func(ctx echo.Context) error {
		data := "Hello World"
		return ctx.String(http.StatusOK, data)
	})

	route.GET("/html", func(ctx echo.Context) error {
		data := "Hello /html"
		return ctx.HTML(http.StatusOK, data)
	})

	route.GET("/redirect", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	route.GET("/json", func(ctx echo.Context) error {
		data := M{
			"Name": "Reyhan Alphard Savero",
			"Type": "Json",
			"Job":  "Backend Engineer",
		}

		return ctx.JSON(http.StatusOK, data)
	})

	route.Start(":8080")
}

package main

import (
	"net/http"

	"github.com/labstack/echo"
)

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

	route.Start(":8080")
}

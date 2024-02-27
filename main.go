package main

import (
	"context"

	"github.com/Chasmyr/echo-app/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// main menu
	component := templates.Hello("John")

	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":3000"))
}

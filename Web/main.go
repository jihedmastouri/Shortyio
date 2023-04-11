package main

import (
	"github.com/labstack/echo/v4"
	"github.com/shorty-io/go-shorty/web/config"
	"github.com/shorty-io/go-shorty/web/handler"
)

func main() {

	e := echo.New()

	conn := config.Connect()
    defer conn.Close()

	e.GET("/", func(c echo.Context) error {
		return handler.CallService(c, conn)
	})

	e.Logger.Fatal(e.Start("localhost:8080"))
}

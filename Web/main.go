package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shorty-io/go-shorty/Shared/service"
	"github.com/shorty-io/go-shorty/web/handler"
)

func main() {
	srv := service.New("Web")

    // Not necessary at the moment
	c := service.InitConfig{
		ServiceRegister: service.Consul,
		ConfigProvider: service.ConsulConfig,
	}
	srv.Init(c)

	srv.Start()

	e := echo.New()

    e.GET("/:lang/:id", func(c echo.Context) error {
		e.Logger.Info("Start")
		conn, err := srv.Dial("Queries", "")
		if err != nil {
			e.Logger.Fatal(err)
		}
		return handler.CallService(c, conn)
	})

    e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

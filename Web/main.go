package main

import (
	"github.com/labstack/echo/v4"
	"github.com/shorty-io/go-shorty/Shared/service"

	"github.com/shorty-io/go-shorty/web/handler"
	"github.com/shorty-io/go-shorty/web/handler/flipflop"
	"github.com/shorty-io/go-shorty/web/handler/queries"
)

func main() {
	srv := service.New("Web")
	srv.Start()
	defer handler.Cleanup()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {

		// "BlockUpdated", "BlockUpdatedQ"
		// nc.Publish("BlockUpdated", []byte("Hello World"))
		return c.String(200, "Hello, World!")
	})

	queries.New(e, srv.Dial)
	flipflop.New(e, srv.Dial)

	e.Logger.Fatal(e.Start(":8080"))
}

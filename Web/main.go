package main

import (

	"github.com/labstack/echo/v4"
	"github.com/shorty-io/go-shorty/Shared/service"
	"github.com/shorty-io/go-shorty/web/config"
	"github.com/shorty-io/go-shorty/web/handler"
)

func main() {
	srv := service.New("Web")
	srv.Start()

	e := echo.New()

	m := config.NewMicroS()

	// Get Block Metadata and list of languages
	e.GET("/", func(c echo.Context) error {
		return c.String(200,"Hello world!")
	})

	// Get Block Metadata and content for a language
    e.GET("/get-block/:lang/:id", func(c echo.Context) error {
		return handler.GetBlock(c, m.Queries)
	})

	// Get All Versions
    e.GET("/versions/:lang/:id/version", func(c echo.Context) error {
		return handler.GetVersions(c, m.Queries)
	})

	// Get All Versions
    e.GET("/langs/:id", func(c echo.Context) error {
		return handler.GetLanguages(c, m.Queries)
	})

	//
	// // Create a new Language for a Block
	// e.POST("/newBlockLang", func(c echo.Context) error {
	// 	return handler.GetBlock(c, &m.FlipFlop)
	// })
	//
	// // Delete a Whole Block
	// e.DELETE("/:id", func(c echo.Context) error {
	// 	return handler.GetBlock(c, &m.FlipFlop)
	// })
	//
	// // Delete a Block language
	// e.DELETE("/:lang/:id", func(c echo.Context) error {
	// 	return handler.GetBlock(c, &m.FlipFlop)
	// })

	e.Logger.Fatal(e.Start(":8080"))
}

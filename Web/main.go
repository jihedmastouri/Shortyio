package main

import (
	"github.com/labstack/echo/v4"
	"github.com/shorty-io/go-shorty/Shared/service"
	q "github.com/shorty-io/go-shorty/queries/proto"
	"github.com/shorty-io/go-shorty/web/handler"
)

func main() {
	srv := service.New("Web")
	srv.Start()

	e := echo.New()

	nc, err := nats.Connect("nats://localhost:4222")
    if err != nil {
        log.Fatal(err)
    }
    defer nc.Close()

	// Block/full/Lang/id
	// Block/meta/lang/id
	// Block/content/lang/id
	// Block/rules/lang/id
	// Block/languages/id

	e.GET("/", func(c echo.Context) error {
		// "BlockUpdated", "BlockUpdatedQ"
		return c.String(200, "Hello world!")
	})

	block := e.Group("/block")

	// Get Block Metadata and content for a language
	block.GET("/full/:lang/:id", func(c echo.Context) error {
		connQuery, err := srv.Dial(service.Queries, nil)
		if err != nil {
			c.Logger().Debug(err)
		}

		defer connQuery.Close()

		clientQuery := q.NewQueriesClient(connQuery)
		return handler.GetBlock(c, clientQuery)
	})

	// Get All Versions
	block.GET("/versions/:lang/:id", func(c echo.Context) error {
		connQuery, err := srv.Dial(service.Queries, nil)
		if err != nil {
			c.Logger().Debug(err)
		}

		defer connQuery.Close()

		clientQuery := q.NewQueriesClient(connQuery)
		return handler.GetVersions(c, clientQuery)
	})

	// Get All Versions
	block.GET("/languages/:id", func(c echo.Context) error {

		connQuery, err := srv.Dial(service.Queries, nil)
		if err != nil {
			c.Logger().Debug(err)
		}

		defer connQuery.Close()

		clientQuery := q.NewQueriesClient(connQuery)
		return handler.GetLanguages(c, clientQuery)
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

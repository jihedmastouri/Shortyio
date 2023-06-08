package main

import (

	"github.com/labstack/echo/v4"
	"github.com/shorty-io/go-shorty/Shared/service"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"github.com/shorty-io/go-shorty/web/handler"
)

func main() {
	srv := service.New("Web")
	srv.Start()
	defer handler.Cleanup()

	e := echo.New()


	// Block/full/Lang/id
	// Block/meta/lang/id
	// Block/content/lang/id
	// Block/rules/lang/id
	// Block/languages/id

	e.GET("/", func(c echo.Context) error {

		// "BlockUpdated", "BlockUpdatedQ"
		// nc.Publish("BlockUpdated", []byte("Hello World"))
		return c.String(200, "Hello, World!")
	})

	block := e.Group("/block")
	Commands := block.Group("/Commands")

	// Get Block Metadata and content for a language
	block.GET("/full/:lang/:id", func(c echo.Context) error {
		connQuery, err := srv.Dial(service.Queries, nil)
		if err != nil {
			c.Logger().Debug(err)
		}

		defer connQuery.Close()

		clientQuery := pb.NewQueriesClient(connQuery)
		return handler.GetBlock(c, clientQuery)
	})

	// Get All Versions
	block.GET("/versions/:lang/:id", func(c echo.Context) error {
		connQuery, err := srv.Dial(service.Queries, nil)
		if err != nil {
			c.Logger().Debug(err)
		}

		defer connQuery.Close()

		clientQuery := pb.NewQueriesClient(connQuery)
		return handler.GetVersions(c, clientQuery)
	})

	// Get All Versions
	block.GET("/languages/:id", func(c echo.Context) error {

		connQuery, err := srv.Dial(service.Queries, nil)
		if err != nil {
			c.Logger().Debug(err)
		}

		defer connQuery.Close()

		clientQuery := pb.NewQueriesClient(connQuery)
		return handler.GetLanguages(c, clientQuery)
	})

	// Create a new Block
	Commands.POST("/new", func(c echo.Context) error {
		connCommand, err := srv.Dial(service.FlipFlop, nil)
		if err != nil {
			c.Logger().Debug(err)
		}

		defer connCommand.Close()

		clientCommand := pb.NewFlipFlopClient(connCommand)
		return handler.CreateBlock(c, clientCommand)
	})

	// Update Content
	Commands.POST("/content", func(c echo.Context) error {
		return handler.UpdateContent(c)
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

package flipflop

import (
	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"github.com/shorty-io/go-shorty/Shared/service/namespace"

	"github.com/shorty-io/go-shorty/web/handler"
)

func New(e *echo.Echo, fn handler.Dialfn) {
	CreateClient := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conn, err := fn(namespace.FlipFlop, nil)
			if err != nil {
				c.Logger().Debug(err)
				return err
			}
			defer conn.Close()
			client := pb.NewFlipFlopClient(conn)

			c.Set("client", client)

			err = next(c)

			return err
		}
	}

	cmd := e.Group("/cmd", CreateClient)

	cmd.POST("/block", createBlock)
	cmd.DELETE("/block/:id", deleteBlock)
	cmd.PUT("/block/:id", updateBlock)

	cmd.POST("/block/:id/lang", createLanguage)
	cmd.DELETE("/block/:id/lang/:code", deleteLanguage)

	cmd.POST("/block-rule", createRule)
	cmd.DELETE("/block/:id", deleteRule)
	cmd.PUT("/block/:id", updateRule)

	cmd.POST("/tag", createTag)
	cmd.DELETE("/tag/:id", deleteTag)
	cmd.POST("/block/:id/tag/:taxonomy-id", addTagToBlock)

	cmd.POST("/categ", createCateg)
	cmd.DELETE("/categ/:id", deleteCateg)
	cmd.POST("/block/:id/categ/:taxonomy-id", addCategToBlock)
}

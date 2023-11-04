package flipflop

import (
	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"github.com/shorty-io/go-shorty/Shared/service/namespace"
	"google.golang.org/grpc"
)

type Dialfn func(serviceName namespace.DefaultServices, tag *[]string) (*grpc.ClientConn, error)

func New(e *echo.Echo, fn Dialfn, authMiddleware echo.MiddlewareFunc) {

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

	cmd := e.Group("/cmd",
		CreateClient,
		authMiddleware,
	)

	// Blocks
	cmd.POST("/block", createBlock)
	cmd.DELETE("/block/:id", deleteBlock)
	cmd.PUT("/block/:id", updateBlock)

	// Languages
	cmd.POST("/block/:id/lang", createLanguage)
	cmd.DELETE("/block/:id/lang/:code", deleteLanguage)

	// Rules
	cmd.POST("/block-rule", createRule)
	cmd.DELETE("/block-rule/:name", deleteRule)
	cmd.PUT("/block-rule/:name", updateRule)

	// Tags
	cmd.POST("/tag", createTag)
	cmd.DELETE("/tag/:name", deleteTag)
	cmd.POST("/block/:id/tag/:names", addTagToBlock)

	// Categories
	cmd.POST("/categ", createCateg)
	cmd.DELETE("/categ/:name", deleteCateg)
	cmd.POST("/block/:id/categ/:names", addCategToBlock)
}

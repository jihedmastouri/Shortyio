package queries

import (
	"github.com/labstack/echo/v4"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"github.com/shorty-io/go-shorty/Shared/service/namespace"
	"github.com/shorty-io/go-shorty/web/handler"
)

func New(e *echo.Echo, fn handler.Dialfn) {

	createClient := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conn, err := fn(namespace.Queries, nil)
			if err != nil {
				c.Logger().Debug(err)
				return err
			}
			defer conn.Close()
			client := pb.NewQueriesClient(conn)

			c.Set("client", client)

			err = next(c)

			return err
		}
	}

	block := e.Group("/public/block", createClient)

	block.GET("/full/:lang/:id", getBlock)
	block.GET("/content/:lang/:id", getBlockContent)
	block.GET("/meta/:lang/:id", getBlockMeta)
	block.GET("/rules/:lang/:id", getBlockRules)

	block.GET("/versions/:lang/:id", getVersions)
	block.GET("/languages/:id", getLanguages)

	block.GET("/search/:lang", searchBlock)
}

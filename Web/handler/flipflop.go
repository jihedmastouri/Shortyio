package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/queries/proto"
	"google.golang.org/grpc"
)

func CallService(c echo.Context, conn *grpc.ClientConn) error {
	client := pb.NewQueriesClient(conn)
	req := &pb.BlockRequest{Id: c.Param("id"), Lang: c.Param("lang")}
	e := c.Echo()
	e.Logger.Info("Done")

	res, err := client.GetBlock(context.Background(), req)
	if err != nil {
		e.Logger.Error(err)
		log.Print(err)
        return c.JSON(http.StatusNotFound, echo.Map{ "err": err })
	}

	e.Logger.Info(res)
	return c.JSON(http.StatusOK, res)
}

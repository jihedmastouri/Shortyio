package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/queries/proto"
)

func GetBlock(c echo.Context, client pb.QueriesClient) error {
	req := &pb.BlockRequest{Id: c.Param("id"), Lang: c.Param("lang")}
    log.Println("GetBlock", "Starting")

	res, err := client.GetBlock(context.Background(), req)
	if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{ "err": err })
	}

	return c.JSON(http.StatusOK, res)
}

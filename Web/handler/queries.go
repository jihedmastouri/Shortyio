package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/queries/proto"
)

func GetBlock(c echo.Context, client pb.QueriesClient) error {

	req := &pb.BlockRequest{Id: c.Param("id"), Lang: c.Param("lang")}

	versionStr := c.QueryParam("version")
	version, err := strconv.Atoi(versionStr)
	if err == nil {
		log.Println(err)
        temp := int32(version)
		req.Version = &temp
	}

	log.Println("GetBlock", "Starting")

	res, err := client.GetBlock(context.Background(), req)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound, echo.Map{"err": err})
	}

	return c.JSON(http.StatusOK, res)
}

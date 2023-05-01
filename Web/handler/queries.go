package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/queries/proto"
	"google.golang.org/grpc"
)

func GetBlock(c echo.Context, client pb.QueriesClient) error {
	return getStuffBlock(client.GetBlock, c)
}

func GetVersions(c echo.Context, client pb.QueriesClient) error {
	req := &pb.VersionsRequest{Id: c.Param("id"), Lang: c.Param("lang")}

	log.Println("GetVersion", "Starting")
	res, err := client.GetVersions(context.Background(), req)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func GetLanguages(c echo.Context, client pb.QueriesClient) error {
	req := &pb.LanguageRequest{Id: c.Param("id")}

	log.Println("GetLanguages", "Starting")
	res, err := client.GetLanguages(context.Background(), req)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

type blockStuff interface {
	pb.Block | pb.BlockMeta | pb.BlockRules | pb.BlockContent
}

type fnc[T blockStuff] func(context.Context, *pb.BlockRequest, ...grpc.CallOption) (*T, error)

func getStuffBlock[T blockStuff](fn fnc[T], c echo.Context) error {
	req := &pb.BlockRequest{Id: c.Param("id"), Lang: c.Param("lang")}

	versionStr := c.QueryParam("version")
	version, err := strconv.Atoi(versionStr)
	if err == nil {
		log.Println(err)
		temp := int32(version)
		req.Version = &temp
	}

	log.Println("GetBlock", "Starting")
	res, err := fn(context.Background(), req)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

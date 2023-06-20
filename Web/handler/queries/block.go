package queries

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"google.golang.org/grpc"
)

func getBlock(c echo.Context) error {
	client := c.Get("client").(pb.QueriesClient)
	return getStuffBlock(client.GetBlock, c)
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

package queries

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var marshaller = protojson.MarshalOptions{
	Multiline:       true,
	Indent:          "  ",
	AllowPartial:    true,
	UseProtoNames:   false,
	UseEnumNumbers:  false,
	EmitUnpopulated: true,
}

// Get Block Like Crazy
//
//	@Summary      Get Block
//	@Description  get block by id and lang
//	@Produce      json
//	@Param        id    path     string
//	@Param        lang    path   string
//	@Success      200  {array}   pb.Block
//	@Failure      400  {object}  httputil.HTTPError
//	@Failure      404  {object}  httputil.HTTPError
//	@Failure      500  {object}  httputil.HTTPError
//	@Router       /full [get]
func getBlock(c echo.Context) error {
	client := c.Get("client").(pb.QueriesClient)
	res, err := getStuffBlock(client.GetBlock, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	json, err := marshaller.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSONBlob(http.StatusOK, json)

}

func getBlockContent(c echo.Context) error {
	client := c.Get("client").(pb.QueriesClient)
	res, err := getStuffBlock(client.GetBlockContent, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	json, err := marshaller.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSONBlob(http.StatusOK, json)
}

func getBlockMeta(c echo.Context) error {
	client := c.Get("client").(pb.QueriesClient)
	res, err := getStuffBlock(client.GetBlockMeta, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	json, err := marshaller.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSONBlob(http.StatusOK, json)
}

func getBlockRules(c echo.Context) error {
	client := c.Get("client").(pb.QueriesClient)
	res, err := getStuffBlock(client.GetBlockRules, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	json, err := marshaller.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSONBlob(http.StatusOK, json)
}

type blockStuff interface {
	pb.Block | pb.BlockMeta | pb.BlockRules | pb.BlockContent
}

type fnc[T blockStuff] func(context.Context, *pb.BlockRequest, ...grpc.CallOption) (*T, error)

func getStuffBlock[T blockStuff](fn fnc[T], c echo.Context) (*T, error) {
	req := &pb.BlockRequest{Id: c.Param("id"), Lang: c.Param("lang")}

	versionStr := c.QueryParam("version")
	version, err := strconv.Atoi(versionStr)
	if err == nil {
		log.Println(err)
		temp := int32(version)
		req.Version = &temp
	}

	return fn(context.Background(), req)
}

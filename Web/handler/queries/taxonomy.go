package queries

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

func getAllTags(c echo.Context) error {
	client := c.Get("client").(pb.QueriesClient)
	res, err := client.GetTags(context.Background(), &pb.TagListRq{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	json, err := marshaller.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSONBlob(http.StatusOK, json)
}

func getAllCategories(c echo.Context) error {
	client := c.Get("client").(pb.QueriesClient)
	res, err := client.GetCategories(context.Background(), &pb.CategoryListRq{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	json, err := marshaller.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSONBlob(http.StatusOK, json)
}

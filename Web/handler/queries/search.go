package queries

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

func searchBlock(c echo.Context) error {
	client := c.Get("client").(pb.QueriesClient)
	l := c.Logger()

	tagsParam := c.QueryParam("tags")
	var tags []string
	if tagsParam != "" {
		tags = strings.Split(tagsParam, ",")
	}

	catgsParam := c.QueryParam("catgs")
	var catgs []string
	if catgsParam != "" {
		catgs = strings.Split(catgsParam, ",")
	}

	typeName := c.QueryParam("type")

	pageSizeParam := c.QueryParam("pagesize")
	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		l.Warn(err)
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 100
	}

	pageNumParam := c.QueryParam("pagenum")
	pageNum, err := strconv.Atoi(pageNumParam)
	if err != nil {
		l.Warn(err)
	}
	if pageNum <= 0 {
		pageNum = 1
	}

	selectors := &pb.Selectors{
		Tags:       tags,
		Categories: catgs,
		Type:       typeName,
	}

	pagination := &pb.Pagination{
		PageSize: uint32(pageSize),
		PageNum:  uint32(pageNum),
	}

	res, err := client.Search(context.Background(), &pb.SearchRequest{
		Selectors:  selectors,
		Pagination: pagination,
	})
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	json, err := marshaller.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSONBlob(http.StatusOK, json)
}

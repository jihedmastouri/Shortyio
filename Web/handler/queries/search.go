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

	tagsParam := c.QueryParam("tags")
	tags := strings.Split(tagsParam, ",")

	catgsParam := c.QueryParam("catgs")
	catgs := strings.Split(catgsParam, ",")

	typeName := c.QueryParam("type")

	l := c.Logger()

	pageSizeParam := c.QueryParam("pagesize")
	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		l.Warn(err)
	}

	pageNumParam := c.QueryParam("pagenum")
	pageNum, err := strconv.Atoi(pageNumParam)
	if err != nil {
		l.Warn(err)
	}

	selectors := &pb.Selectors{
		Tags:       tags,
		Categories: catgs,
		Type:       typeName,
	}

	pagination := &pb.Pagination{
		PageSize: int32(pageSize),
		PageNum:  int32(pageNum),
	}

	res, err := client.Search(context.Background(), &pb.SearchRequest{
		Selectors:  selectors,
		Pagination: pagination,
	})
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

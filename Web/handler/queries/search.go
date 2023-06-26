package queries

import (
	"context"
	"log"
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

	log.Println(tags)

	catgsParam := c.QueryParam("catgs")
	catgs := strings.Split(catgsParam, ",")

	log.Println(catgs)

	typeName := c.QueryParam("type")

	log.Println(catgs)

	l := c.Logger()

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

	return c.JSON(http.StatusOK, res)
}

package flipflop

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

// createTag func
func createTag(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	var trq TaxonomyRq
	if err := c.Bind(&trq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := &pb.CreateTaxonomy{
		Name:  trq.Name,
		Descr: trq.Descr,
	}

	res, err := client.CreateTag(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// deleteTag func
func deleteTag(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	name := c.Param("name")

	req := &pb.DeleteTaxonomy{
		Name: name,
	}

	res, err := client.DeleteTag(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// addTagToBlock func
func addTagToBlock(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	blockID := c.Param("id")
	taxParam := c.QueryParam("names")
	var taxNames []string
	if taxParam != "" {
		taxNames = strings.Split(taxParam, ",")
	}

	req := &pb.JoinTaxonomy{
		BlockId: blockID,
		Names:   taxNames,
	}

	res, err := client.JoinTag(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// createCateg func
func createCateg(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	var trq TaxonomyRq
	if err := c.Bind(&trq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := &pb.CreateTaxonomy{
		Name:  trq.Name,
		Descr: trq.Descr,
	}

	res, err := client.CreateCategory(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// deleteCateg func
func deleteCateg(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	name := c.Param("name")

	req := &pb.DeleteTaxonomy{
		Name: name,
	}

	res, err := client.DeleteCategory(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// addCategToBlock func
func addCategToBlock(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	blockID := c.Param("id")
	taxParam := c.QueryParam("names")
	var taxNames []string
	if taxParam != "" {
		taxNames = strings.Split(taxParam, ",")
	}

	req := &pb.JoinTaxonomy{
		BlockId: blockID,
		Names:   taxNames,
	}

	res, err := client.JoinCategory(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

package queries

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

func getVersions(c echo.Context) error {
	client := c.Get("client").(pb.QueriesClient)
	req := &pb.VersionsRequest{Id: c.Param("id"), Lang: c.Param("lang")}

	log.Println("GetVersion", "Starting")
	res, err := client.GetVersions(context.Background(), req)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	json, err := marshaller.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSONBlob(http.StatusOK, json)
}

func getLanguages(c echo.Context) error {
	client := c.Get("client").(pb.QueriesClient)
	req := &pb.LanguageRequest{Id: c.Param("id")}

	log.Println("GetLanguages", "Starting")
	res, err := client.GetLanguages(context.Background(), req)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	json, err := marshaller.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"err": err.Error()})
	}

	return c.JSONBlob(http.StatusOK, json)
}

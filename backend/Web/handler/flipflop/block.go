package flipflop

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

func createBlock(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	var brq BlockRq
	if err := c.Bind(&brq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := &pb.CreateRequest{
		Name:        brq.Name,
		BlockType:   brq.BlockType,
		Author:      brq.Author,
		Description: brq.Description,
		Rules: &pb.BlockRulesRq{
			BlockRules: &pb.BlockRulesRq_Rules{
				Rules: &pb.BlockRules{
					RuleName:          brq.Rules.RuleName,
					Nested:            brq.Rules.Nested,
					HasLikes:          brq.Rules.HasLikes,
					HasComments:       brq.Rules.HasComments,
					CommentsHasLikes:  brq.Rules.CommentsHasLikes,
					CommentsEditable:  brq.Rules.CommentsEditable,
					CommentsMaxNested: int32(brq.Rules.CommentsMaxNested),
				},
			},
		},
	}

	res, err := client.CreateBlock(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func deleteBlock(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	id := c.Param("id")

	req := &pb.DeleteRequest{
		Id: id,
	}

	res, err := client.DeleteBlock(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func createLanguage(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	id := c.Param("id")

	var langRq LangRq
	if err := c.Bind(&langRq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := &pb.CreateLangRequest{
		BlockId:      id,
		LangName:     langRq.Name,
		LangCode:     langRq.Code,
		PreviousLang: langRq.PreviousLang,
	}

	res, err := client.CreateBlockLang(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func deleteLanguage(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	id := c.Param("id")
	code := c.Param("code")

	req := &pb.DeleteLangRequest{
		Id:       id,
		LangCode: code,
	}

	res, err := client.DeleteBlockLang(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

package flipflop

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

func updateBlock(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	var brq BlockRq
	if err := c.Bind(&brq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")

	req := &pb.CreateRequest{
		Id:        &id,
		Name:      brq.Name,
		BlockType: brq.BlockType,
		Author:    brq.Author,
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

	res, err := client.UpdateBlock(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func updateBlockMeta(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	var brq BlockMetaRq
	if err := c.Bind(&brq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")

	req := &pb.BlockMeta{
		BlockId:    id,
		Name:       brq.Name,
		Tags:       brq.Tags,
		Categories: brq.Categories,
	}

	res, err := client.UpdateBlocMeta(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func updateBlockRules(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	var brq BlockRulesRq
	if err := c.Bind(&brq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")

	req := &pb.BlockRulesRq{
		BlockId: id,
		BlockRules: &pb.BlockRulesRq_Rules{
			Rules: &pb.BlockRules{
				RuleName:          brq.RuleName,
				Nested:            brq.Nested,
				HasLikes:          brq.HasLikes,
				HasComments:       brq.HasComments,
				CommentsHasLikes:  brq.CommentsHasLikes,
				CommentsEditable:  brq.CommentsEditable,
				CommentsMaxNested: int32(brq.CommentsMaxNested),
			},
		},
	}

	res, err := client.UpdateBlockRule(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

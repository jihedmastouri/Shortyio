package flipflop

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

func createRule(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	var rrq RuleRq
	if err := c.Bind(&rrq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := &pb.BlockRulesRq{
		BlockRules: &pb.BlockRulesRq_Rules{
			Rules: &pb.BlockRules{
				RuleName:          rrq.RuleName,
				Nested:            rrq.Nested,
				HasLikes:          rrq.HasLikes,
				HasComments:       rrq.HasComments,
				CommentsNested:    rrq.CommentsNested,
				CommentsHasLikes:  rrq.CommentsHasLikes,
				CommentsEditable:  rrq.CommentsEditable,
				CommentsMaxNested: int32(rrq.CommentsMaxNested),
				Descr:             rrq.Descr,
			},
		},
	}

	res, err := client.CreateBlockRule(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func deleteRule(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	name := c.Param("name")
	c.Logger().Debug(name)

	req := &pb.BlockRulesRq{
		BlockRules: &pb.BlockRulesRq_RuleName{
			RuleName: name,
		},
	}

	res, err := client.DeleteBlockRule(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func updateRule(c echo.Context) error {
	client := c.Get("client").(pb.FlipFlopClient)

	name := c.Param("name")
	c.Logger().Debug(name)

	var rrq RuleRq
	if err := c.Bind(&rrq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := &pb.BlockRulesRq{
		BlockRules: &pb.BlockRulesRq_Rules{
			Rules: &pb.BlockRules{
				RuleName:          name,
				Nested:            rrq.Nested,
				HasLikes:          rrq.HasLikes,
				HasComments:       rrq.HasComments,
				CommentsNested:    rrq.CommentsNested,
				CommentsHasLikes:  rrq.CommentsHasLikes,
				CommentsEditable:  rrq.CommentsEditable,
				CommentsMaxNested: int32(rrq.CommentsMaxNested),
				Descr:             rrq.Descr,
			},
		},
	}

	res, err := client.UpdateBlockRule(context.Background(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

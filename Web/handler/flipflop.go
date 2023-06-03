package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/flipFlop/proto"
)

type CreateBlockRq struct {
	Name      string `json:"name"`
	BlockType string `json:"blockType"`
	Author    string `json:"author"`
	Rules     struct {
		RuleName          string `json:"ruleName"`
		Nested            bool   `json:"nested"`
		HasLikes          bool   `json:"hasLikes"`
		HasComments       bool   `json:"hasComments"`
		CommentsNested    bool   `json:"commentsNested"`
		CommentsHasLikes  bool   `json:"commentsHasLikes"`
		CommentsEditable  bool   `json:"commentsEditable"`
		CommentsMaxNested int    `json:"commentsMaxNested"`
	}
}

func CreateBlock(c echo.Context, client pb.FlipFlopClient) error {
	var brq CreateBlockRq
	if err := c.Bind(brq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := &pb.CreateRequest{
		Name:      brq.Name,
		BlockType: brq.BlockType,
		Author:    brq.Author,
		Rules: &pb.BlockRules{
			BlockRules: &pb.BlockRules_Rules_{
				Rules: &pb.BlockRules_Rules{
					RuleName:          brq.Rules.RuleName,
					Nested:            brq.Rules.Nested,
					HasLikes:          brq.Rules.HasLikes,
					HasComments:       brq.Rules.HasComments,
					CommentsNested:    brq.Rules.CommentsNested,
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

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	ps "github.com/shorty-io/go-shorty/Shared/nats"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

type UpdateBlockRq struct {
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

func UpdateContent(c echo.Context) error {
	var brq CreateBlockRq
	if err := c.Bind(&brq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	a := &ps.BlockAdded{
		BlockMeta: &pb.BlockMeta{
			BlockId:    "",
			Name:       "",
			Type:       "",
			Tags:       []string{},
			Categories: []string{},
			Authors:    []*pb.Author{},
			UpdatedAt:  "",
			CreatedAt:  "",
		},
		BlockRules: &pb.BlockRules{
			BlockRules: nil,
		},
		Event: ps.Event{},
	}

	out, err := json.Marshal(a)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	nc.Publish(
		"blockToUpdate",
		out,
	)

	return c.JSON(http.StatusOK, "Block updating...")
}

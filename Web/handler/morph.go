package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	mq "github.com/shorty-io/go-shorty/Shared/msgq"
	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

func UpdateContent(c echo.Context) error {
	var brq *pb.BlockContent
	if err := c.Bind(brq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")

	ev := mq.NewEvent("blockToUpdate", mq.BlockToUpdateData{
		Id:           id,
		BlockContent: brq,
	})

	out, err := json.Marshal(ev)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	nc.Publish(
		string(ev.Name),
		out,
	)

	return c.JSON(http.StatusOK, "Block updating...")
}

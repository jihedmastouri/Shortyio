package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
)

type MsgU struct {
	Id        string `json:"id"`
	LangCode  string `json:"lang_code"`
	Content   string `json:"content"`
	ChangeLog string `json:"change_log"`
}

func UpdateContent(c echo.Context, nc *nats.Conn) error {
	var brq MsgU
	if err := c.Bind(brq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	out, err := json.Marshal(brq)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	nc.Publish(
		"BlockUpdate",
		out,
	)

	return c.JSON(http.StatusOK, "Block updating...")
}

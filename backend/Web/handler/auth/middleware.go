package auth

import (
	"context"

	"github.com/labstack/echo/v4"
)

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.NewHTTPError(401, "Unauthorized")
		}

		ctx := context.Background()
		res, err := keyCloak.Conn.RetrospectToken(ctx, token, keyCloak.ClientId, keyCloak.Secret, keyCloak.Realm)
		if err != nil {
			return echo.NewHTTPError(401, "Unauthorized")
		}

		if !*res.Active {
			return echo.NewHTTPError(401, "Unauthorized")
		}

		err = next(c)
		return err
	}
}

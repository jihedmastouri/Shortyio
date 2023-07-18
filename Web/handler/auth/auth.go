package auth

import (
	"context"

	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) {
	e.GET("/login", login)
	e.GET("/logout", logout)
}

func login(c echo.Context) error {
	rq := new(LoginRq)
	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	ctx := context.Background()
	jwt, err := keyCloak.Conn.Login(
		ctx,
		keyCloak.ClientId,
		keyCloak.Secret,
		keyCloak.Realm,
		rq.Username,
		rq.Password,
	)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, LoginRs{
		Token:        jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	})
}

func logout(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	ctx := context.Background()
	err := keyCloak.Conn.Logout(ctx, token, keyCloak.ClientId, keyCloak.Secret, keyCloak.Realm)
	if err != nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	return c.JSON(200, "Logout success")
}

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

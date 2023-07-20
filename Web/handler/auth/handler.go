package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) {
	e.GET("/login", Login)
	e.GET("/refresh", Refresh)
	e.GET("/logout", Logout)
}

func Login(c echo.Context) error {
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

	month := time.Now().Add(30 * 24 * 3600 * time.Second) // 30 days

	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    jwt.RefreshToken,
		Expires:  month,
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	}

	c.SetCookie(&cookie)

	return c.JSON(200, LoginRs{
		Token:        jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	})
}

func Refresh(c echo.Context) error {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	ctx := context.Background()
	jwt, err := keyCloak.Conn.RefreshToken(
		ctx,
		keyCloak.ClientId,
		keyCloak.Secret,
		keyCloak.Realm,
		refreshToken.Value,
	)
	if err != nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	month := time.Now().Add(30 * 24 * 3600 * time.Second) // 30 days

	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    jwt.RefreshToken,
		Expires:  month,
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	}

	c.SetCookie(&cookie)

	return c.JSON(200, LoginRs{
		Token:        jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	})
}

func Logout(c echo.Context) error {
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

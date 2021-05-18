package inhouse_middleware

import (
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	"strings"
)

func LineJwtCheckMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqToken := c.Request().Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer")

			if len(splitToken) != 2 {
				return echo.NewHTTPError(400, "Bad Request")
			}

			data := url.Values{}
			data.Set("id_token", strings.TrimSpace(splitToken[1]))
			data.Set("client_id", "1653759706")

			res, err := http.Post( "https://api.line.me/oauth2/v2.1/verify", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))

			if err != nil || res.StatusCode != 200 {
				return echo.NewHTTPError(401, "Invalid Token")
			}

			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}
package inhouse_middleware

import (
	"context"
	"github.com/labstack/echo"
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/firebase_admin"
	"strings"
)

type FirebaseTokenVerifiedMiddleware interface {
	GetFirebaseTokenVerifiedMiddleware() echo.MiddlewareFunc
}

type firebaseTokenVerifiedMiddleware struct {
	fireBaseAdminClient firebase_admin.Client
}

func FirebaseTokenVerifiedMiddlewareImpl(fireBaseAdminClient firebase_admin.Client) FirebaseTokenVerifiedMiddleware {
	return &firebaseTokenVerifiedMiddleware{fireBaseAdminClient}
}

func (m *firebaseTokenVerifiedMiddleware) GetFirebaseTokenVerifiedMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqToken := c.Request().Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer")

			if len(splitToken) != 2 {
				return echo.NewHTTPError(400, "Bad Request")
			}

			client := m.fireBaseAdminClient.GetFirebaseAdminClient()
			_, err := client.VerifyIDToken(context.Background(), strings.TrimSpace(splitToken[1]))

			if err != nil {
				return echo.NewHTTPError(401, "UnAuthorize")
			}

			if err := next(c); err != nil {
				c.Error(err)
			}

			return nil
		}
	}
}

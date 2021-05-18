package inhouse_middleware

import (
	"github.com/labstack/echo"
)

type firebaseTokenPathVerifiedMiddleware struct {
}

func FirebaseTokenVerifiedPassMiddlewareImpl() FirebaseTokenVerifiedMiddleware {
	return &firebaseTokenPathVerifiedMiddleware{}
}

func (m *firebaseTokenPathVerifiedMiddleware) GetFirebaseTokenVerifiedMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			return err
		}
	}
}
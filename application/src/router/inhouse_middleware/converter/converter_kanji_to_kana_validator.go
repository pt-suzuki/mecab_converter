package converter

import (
	"github.com/labstack/echo"
)

type CreateValidatorMiddleware interface {
	GetKanjiToKanaValidatorMiddleware() echo.MiddlewareFunc
}

type createValidatorMiddleware struct {
}

func KanjiToKanaValidatorMiddlewareImpl() CreateValidatorMiddleware {
	return &createValidatorMiddleware{}
}

func (m *createValidatorMiddleware) GetKanjiToKanaValidatorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.QueryParam("word") == "" {
				return echo.NewHTTPError(400, "Bad Request")
			}
			err := next(c)
			return err
		}
	}
}


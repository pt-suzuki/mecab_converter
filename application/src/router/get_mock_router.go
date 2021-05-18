package router

import (
	"github.com/labstack/echo"
	"github.com/pt-suzuki/mecab_converter/application/src/router/inhouse_middleware"
)

func GetMockRouter() *echo.Echo {
	verifiedTokenMiddleware := inhouse_middleware.FirebaseTokenVerifiedPassMiddlewareImpl()
	return RouterImpl(verifiedTokenMiddleware).GetRouter()
}

package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pt-suzuki/mecab_converter/application/src/provider"
	"github.com/pt-suzuki/mecab_converter/application/src/router"
	"github.com/pt-suzuki/mecab_converter/application/src/router/inhouse_middleware"
)

func CreateEcho()*echo.Echo {
	serviceProvider := provider.GetServiceProvider()
	verifiedTokenMiddleware := inhouse_middleware.FirebaseTokenVerifiedMiddlewareImpl(serviceProvider.FireBaseAdminClient)

	echo := router.RouterImpl(verifiedTokenMiddleware).GetRouter()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())
	echo.Use(middleware.CORS())

	return echo
}



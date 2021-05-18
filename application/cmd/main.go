package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/firebase_admin"
	"github.com/pt-suzuki/mecab_converter/application/src/router"
	"github.com/pt-suzuki/mecab_converter/application/src/router/inhouse_middleware"
)

func main() {
	ftv := inhouse_middleware.FirebaseTokenVerifiedMiddlewareImpl(firebase_admin.ClientImpl())
	echo := router.RouterImpl(ftv).GetRouter()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	echo.Start(":8080")
}
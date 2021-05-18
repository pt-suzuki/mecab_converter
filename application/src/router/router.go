package router

import (
	"github.com/labstack/echo"
	"github.com/pt-suzuki/mecab_converter/application/src/provider"
	"github.com/pt-suzuki/mecab_converter/application/src/router/inhouse_middleware"
	"github.com/pt-suzuki/mecab_converter/application/src/router/inhouse_middleware/converter"
)

type Router interface {
	GetRouter() *echo.Echo
}

type router struct {
	firebaseTokenVerifiedMiddleware inhouse_middleware.FirebaseTokenVerifiedMiddleware
}

func RouterImpl(firebaseTokenVerifiedMiddleware inhouse_middleware.FirebaseTokenVerifiedMiddleware) Router {
	return &router{firebaseTokenVerifiedMiddleware}
}

func (m *router) GetRouter() *echo.Echo {
	e := echo.New()

	// serviceProvider := provider.GetServiceProvider()
	actionProvider := provider.GetActionProvider()

	/*
	e.POST("/api/auth/firebase/custom_token/line",
		actionProvider.AuthFirebaseCustomTokenCreateAction.FirebaseCustomTokenCreateActionInvoke(),
		inhouse_middleware.LineJwtCheckMiddleware())

	ftvm := m.firebaseTokenVerifiedMiddleware.GetFirebaseTokenVerifiedMiddleware()
	 */

	api := e.Group("/api")

	api.GET("/converter/kanji/kana", actionProvider.KanjiToKanaAction.KanjiToKanaActionInvoke(), converter.KanjiToKanaValidatorMiddlewareImpl().GetKanjiToKanaValidatorMiddleware())

	return e
}

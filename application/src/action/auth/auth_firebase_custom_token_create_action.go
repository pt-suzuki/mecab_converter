package auth

import (
	"github.com/labstack/echo"
	"github.com/pt-suzuki/mecab_converter/application/src/domain/auth"
	"github.com/pt-suzuki/mecab_converter/application/src/responder"
)

type FirebaseCustomTokenCreateAction interface {
	FirebaseCustomTokenCreateActionInvoke() echo.HandlerFunc
}

type firebaseCustomTokenCreateAction struct {
	service    auth.Service
	translator auth.Translator
}

func FirebaseCustomTokenCreateActionImpl(service auth.Service, translator auth.Translator) FirebaseCustomTokenCreateAction {
	return &firebaseCustomTokenCreateAction{service, translator}
}

func (action *firebaseCustomTokenCreateAction) FirebaseCustomTokenCreateActionInvoke() echo.HandlerFunc {
	return func(context echo.Context) error {
		return responder.AuthFirebaseCustomTokenCreateResponder(context,
			action.service.FireBaseCustomTokenCreate(action.translator.TranslateRequestToAuthId(context)))
	}
}
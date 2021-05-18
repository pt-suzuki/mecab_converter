package provider

import (
	"github.com/pt-suzuki/mecab_converter/application/src/action/auth"
	"github.com/pt-suzuki/mecab_converter/application/src/action/converter"
)

type ActionProvider struct {
	AuthFirebaseCustomTokenCreateAction auth.FirebaseCustomTokenCreateAction
	KanjiToKanaAction        converter.KanjiToKanaAction
}


func GetActionProvider() ActionProvider {
	sp := GetServiceProvider()

	actionProvider := ActionProvider{}

	actionProvider.AuthFirebaseCustomTokenCreateAction = auth.FirebaseCustomTokenCreateActionImpl(sp.AuthService,sp.AuthTranslator)

	actionProvider.KanjiToKanaAction = converter.KanjiToKanaActionImpl(sp.ConverterService, sp.ConverterTranslator)

	return actionProvider
}



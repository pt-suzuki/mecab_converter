package converter

import (
	"github.com/labstack/echo"
	"github.com/pt-suzuki/mecab_converter/application/src/domain/converter"
	"github.com/pt-suzuki/mecab_converter/application/src/responder"
)

type KanjiToKanaAction interface {
	KanjiToKanaActionInvoke() echo.HandlerFunc
}

type kanjiToKanaAction struct {
	service converter.Service
	translator converter.Translator
}

func KanjiToKanaActionImpl(service converter.Service, translator converter.Translator) KanjiToKanaAction {
	return &kanjiToKanaAction{service, translator}
}

func (action *kanjiToKanaAction) KanjiToKanaActionInvoke() echo.HandlerFunc {
	return func(context echo.Context) error {
		result := action.service.ParseKanjiToKana(action.translator.TranslateRequestToWord(context))
		return responder.KanjiToKanaResponder(context, result)
	}
}

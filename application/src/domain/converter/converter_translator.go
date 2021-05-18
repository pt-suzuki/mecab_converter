package converter

import (
	"github.com/bluele/mecab-golang"
	"github.com/labstack/echo"
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/handler"
	"strings"
)

type Translator interface {
	TranslateRequestToWord(context echo.Context) string
	TranslateNodeToString(node *mecab.Node) string
}

type translator struct {
	responseHandler handler.ResponseHandler
}

func TranslatorImpl(responseHandler handler.ResponseHandler) Translator {
	return &translator{ responseHandler}
}

func (t *translator) TranslateRequestToWord(context echo.Context) string {
	word := context.QueryParam("word")

	return word
}

func (t *translator) TranslateNodeToString(node *mecab.Node) string {
	var key string
	for {
		features := strings.Split(node.Feature(), ",")
		key = key + features[7]
		if node.Next() != nil {
			break
		}
	}
	return strings.Trim(key, "*")
}
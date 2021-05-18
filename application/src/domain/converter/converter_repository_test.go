package converter

import (
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/handler"
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/mecab_client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test漢字からカタカナへの変換(t *testing.T) {
	repository := getRepository()

	result := repository.parseKanjiToKana("漢字からカタカナへの変換")
	assert.Equal(t, "カンジカラカタカナヘノヘンカン", result)
}

func getRepository() Repository {
	client := mecab_client.ClientImpl()
	responseHandler := handler.ResponseHandlerImpl()
	translator := TranslatorImpl(responseHandler)

	repository := RepositoryImpl(client, translator)

	return repository
}
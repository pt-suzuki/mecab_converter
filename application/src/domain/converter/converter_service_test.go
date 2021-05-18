package converter

import (
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/handler"
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/mecab_client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Testサービスでユーザーの保存(t *testing.T) {
	service := getService()

	result := service.parseKanjiToKana("漢字からカタカナへの変換")
	assert.Equal(t, "カンジカラカタカナヘノヘンカン", result)
}

func getService() Service {
	client := mecab_client.ClientImpl()
	responseHandler := handler.ResponseHandlerImpl()
	translator := TranslatorImpl(responseHandler)

	repository := RepositoryImpl(client, translator)
	service := ServiceImpl(repository)

	return service
}


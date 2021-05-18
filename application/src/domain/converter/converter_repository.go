package converter

import (
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/mecab_client"
)

type Repository interface {
	ParseKanjiToKana(word string) string
}

type repository struct {
	client mecab_client.Client
	translator Translator
}

func RepositoryImpl(client mecab_client.Client, translator Translator) Repository {
	return &repository{ client,  translator}
}

func (r *repository) ParseKanjiToKana(word string) string {
	client := r.client.GetMecabClient()
	defer client.Destroy()

	tg, err := client.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()

	lt, err := client.NewLattice(word)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)

	return r.translator.TranslateNodeToString(node)
}
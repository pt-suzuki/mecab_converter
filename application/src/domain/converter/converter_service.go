package converter

type Service interface {
	ParseKanjiToKana(word string) string
}

type service struct {
	repository Repository
}

func ServiceImpl(repository Repository) Service {
	return &service{repository}
}

func (s *service) ParseKanjiToKana(word string) string {
	return s.repository.ParseKanjiToKana(word)
}

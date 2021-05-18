package auth

type Service interface {
	FireBaseCustomTokenCreate(Custom string) string
}

type service struct {
	repository Repository
	translator Translator
}

func ServiceImpl(repository Repository, translator Translator) Service {
	return &service{repository, translator}
}

func (s *service) FireBaseCustomTokenCreate(Custom string) string {
	return s.repository.FireBaseCustomTokenCreate(Custom)
}
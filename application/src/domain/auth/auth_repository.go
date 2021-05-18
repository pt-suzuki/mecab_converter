package auth

import (
	"context"
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/firebase_admin"
	"log"
)

type Repository interface {
	FireBaseCustomTokenCreate(id string) string
}

type repository struct {
	translator Translator
	firebaseAdminClient     firebase_admin.Client
}

func RepositoryImpl(translator Translator, firebaseAdminClient firebase_admin.Client) Repository {
	return &repository{translator, firebaseAdminClient}
}

func (r *repository) FireBaseCustomTokenCreate(id string) string {
	client := r.firebaseAdminClient.GetFirebaseAdminClient()
	print(id)
	token, err := client.CustomToken(context.Background(), id)
	if err != nil {
		log.Fatal("error minting custom token: \n", err)
	}
	return token
}



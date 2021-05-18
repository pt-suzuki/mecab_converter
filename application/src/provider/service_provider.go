package provider

import (
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/firebase_admin"
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/firebase_cloud_message"
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/firestore"
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/handler"
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/mecab_client"
	"github.com/pt-suzuki/mecab_converter/application/src/domain/auth"
	"github.com/pt-suzuki/mecab_converter/application/src/domain/converter"
)

type ServiceProvider struct {
	FireStoreClient firestore.Client
	FireBaseAdminClient firebase_admin.Client

	FireBaseCloudMessageClient firebase_cloud_message.Client

	MecabClient mecab_client.Client

	ResponseHandler handler.ResponseHandler

	AuthService auth.Service
	AuthTranslator auth.Translator
	AuthRepository auth.Repository

	ConverterService converter.Service
	ConverterTranslator converter.Translator
	ConverterRepository converter.Repository
}


func GetServiceProvider() ServiceProvider {
	sp := ServiceProvider{}

	sp.FireStoreClient = firestore.ClientImpl()
	sp.FireBaseAdminClient = firebase_admin.ClientImpl()
	sp.FireBaseCloudMessageClient = firebase_cloud_message.ClientImpl()

	sp.MecabClient = mecab_client.ClientImpl()
	sp.ResponseHandler = handler.ResponseHandlerImpl()

	sp.AuthTranslator = auth.TranslatorImpl()
	sp.AuthRepository = auth.RepositoryImpl(sp.AuthTranslator, sp.FireBaseAdminClient)
	sp.AuthService = auth.ServiceImpl(sp.AuthRepository, sp.AuthTranslator)

	sp.ConverterTranslator = converter.TranslatorImpl(sp.ResponseHandler)
	sp.ConverterRepository = converter.RepositoryImpl(sp.MecabClient,sp.ConverterTranslator)
	sp.ConverterService = converter.ServiceImpl(sp.ConverterRepository)

	return sp
}



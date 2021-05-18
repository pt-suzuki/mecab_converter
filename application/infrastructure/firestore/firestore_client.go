package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/pt-suzuki/mecab_converter/application/config"
	"log"
)

type Client interface {
	GetFireStoreClient() *firestore.Client
}

type client struct {
}

func ClientImpl() Client {
	return &client{}
}

func (c *client)GetFireStoreClient () *firestore.Client {
	appConfig := config.NewConfig()

	// Use a service account
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: appConfig.GcpProjectId}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

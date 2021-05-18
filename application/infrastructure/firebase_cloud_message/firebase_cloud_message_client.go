package firebase_cloud_message

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/pt-suzuki/mecab_converter/application/config"
	"log"
)

type Client interface {
	GetMessagingClient() *messaging.Client
}

type client struct {
}

func ClientImpl() Client {
	return &client{}
}

func (c *client)GetMessagingClient () *messaging.Client {
	appConfig := config.NewConfig()

	// Use a service account
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: appConfig.GcpProjectId}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

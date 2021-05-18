package firebase_admin

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"log"
)

type Client interface {
	GetFirebaseAdminClient() *auth.Client
}

type client struct {
}

func ClientImpl() Client {
	return &client{}
}

func (c *client)GetFirebaseAdminClient () *auth.Client {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return client
}
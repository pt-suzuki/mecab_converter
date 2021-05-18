package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"path/filepath"
	"runtime"
)

type serverClient struct {
}

func ServerClientImpl() Client {
	return &serverClient{}
}

func (c *serverClient)GetFireStoreClient () *firestore.Client {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../../")
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(root + "/key/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

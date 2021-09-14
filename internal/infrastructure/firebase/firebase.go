package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/balloon/go/invite/env"
	"google.golang.org/api/option"
)

var client *firestore.Client

func newApp(ctx context.Context) (*firebase.App, error) {
	if env.DEBUG {
		// デバック時にはクレデンシャルファイルから初期化
		opt := option.WithCredentialsFile(env.GoogleApplicationCredentials)
		app, err := firebase.NewApp(ctx, nil, opt)
		return app, err
	} else {
		app, err := firebase.NewApp(ctx, nil)
		return app, err
	}
}

func NewFirestore() (*firestore.Client, error) {
	if client != nil {
		return client, nil
	}

	ctx := context.Background()

	// firebaseを初期化
	app, err := newApp(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while initializing firebase app: %v", err)
	}

	// firebaseを初期化
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while initializing firestore: %v", err)
	}

	return client, nil
}

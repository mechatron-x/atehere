package service

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/mechatron-x/atehere/internal/config"
	"google.golang.org/api/option"
)

type (
	Firebase struct {
		app *firebase.App
	}
)

func NewFirebase(conf config.Firebase) (*Firebase, error) {
	bytes, err := json.Marshal(conf)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsJSON(bytes)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return &Firebase{
		app: app,
	}, nil

}

func (f *Firebase) VerifyUser(idToken string) (*auth.UserRecord, error) {
	client, err := f.app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	authToken, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		return nil, err
	}

	return client.GetUser(context.Background(), authToken.UID)
}

func (f *Firebase) RevokeRefreshTokens(idToken string) error {
	client, err := f.app.Auth(context.Background())
	if err != nil {
		return err
	}

	authToken, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		return err
	}

	return client.RevokeRefreshTokens(context.Background(), authToken.UID)
}

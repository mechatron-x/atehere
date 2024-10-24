package infrastructure

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/mechatron-x/atehere/internal/config"
	"google.golang.org/api/option"
)

type (
	FirebaseAuthenticator struct {
		app *firebase.App
	}
)

func NewFirebaseAuthenticator(conf config.Firebase) (*FirebaseAuthenticator, error) {
	bytes, err := json.Marshal(conf)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsJSON(bytes)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return &FirebaseAuthenticator{
		app: app,
	}, nil

}

func (fa *FirebaseAuthenticator) VerifyUser(idToken string) (*auth.UserRecord, error) {
	client, err := fa.app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	authToken, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		return nil, err
	}

	return client.GetUser(context.Background(), authToken.UID)
}

func (fa *FirebaseAuthenticator) RevokeRefreshTokens(idToken string) error {
	client, err := fa.app.Auth(context.Background())
	if err != nil {
		return err
	}

	authToken, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		return err
	}

	return client.RevokeRefreshTokens(context.Background(), authToken.UID)
}

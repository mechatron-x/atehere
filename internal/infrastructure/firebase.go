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
	FirebaseAuth struct {
		app *firebase.App
	}
)

func NewFirebaseAuth(conf config.Firebase) (*FirebaseAuth, error) {
	bytes, err := json.Marshal(conf)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsJSON(bytes)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return &FirebaseAuth{
		app: app,
	}, nil

}

func (fa *FirebaseAuth) VerifyUser(idToken string) (*auth.UserRecord, error) {
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

func (fa *FirebaseAuth) RevokeRefreshTokens(idToken string) error {
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

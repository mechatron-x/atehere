package authenticator

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

func NewFirebase(conf config.Firebase) (*FirebaseAuthenticator, error) {
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

func (fa *FirebaseAuthenticator) CreateUser(id, email, password string) error {
	client, err := fa.app.Auth(context.Background())
	if err != nil {
		return err
	}

	user := &auth.UserToCreate{}
	user.UID(id)
	user.Email(email)
	user.Password(password)

	_, err = client.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (fa *FirebaseAuthenticator) GetUserID(idToken string) (string, error) {
	record, err := fa.getUserRecord(idToken)
	if err != nil {
		return "", err
	}

	return record.UID, nil
}

func (fa *FirebaseAuthenticator) GetUserEmail(idToken string) (string, error) {
	record, err := fa.getUserRecord(idToken)
	if err != nil {
		return "", err
	}

	return record.Email, nil
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

	err = client.RevokeRefreshTokens(context.Background(), authToken.UID)
	if err != nil {
		return err
	}

	return nil
}

func (fa *FirebaseAuthenticator) getUserRecord(idToken string) (*auth.UserRecord, error) {
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

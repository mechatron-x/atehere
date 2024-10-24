package infrastructure

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
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

	return err
}

func (fa *FirebaseAuthenticator) VerifyUser(idToken string) (*port.AuthRecord, error) {
	client, err := fa.app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	authToken, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		return nil, err
	}

	record, err := client.GetUser(context.Background(), authToken.UID)
	if err != nil {
		return nil, err
	}

	return &port.AuthRecord{
		UID:           record.UID,
		Disabled:      record.Disabled,
		EmailVerified: record.EmailVerified,
		Email:         record.Email,
		PhoneNumber:   record.PhoneNumber,
	}, nil
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

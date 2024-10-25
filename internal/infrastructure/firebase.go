package infrastructure

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/core"
	"google.golang.org/api/option"
)

const (
	pkg string = "infrastructure.Firebase"
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

func (fa *FirebaseAuthenticator) CreateUser(id, email, password string) core.PortError {
	client, err := fa.app.Auth(context.Background())
	if err != nil {
		return core.NewConnectionError(pkg, err)
	}

	user := &auth.UserToCreate{}
	user.UID(id)
	user.Email(email)
	user.Password(password)

	_, err = client.CreateUser(context.Background(), user)
	if err != nil {
		return core.NewDataNotFoundError(pkg, err)
	}

	return nil
}

func (fa *FirebaseAuthenticator) GetUserID(idToken string) (string, core.PortError) {
	record, err := fa.getUserRecord(idToken)
	if err != nil {
		return "", core.NewAuthenticationFailedError(pkg, err)
	}

	return record.UID, nil
}

func (fa *FirebaseAuthenticator) GetUserEmail(idToken string) (string, core.PortError) {
	record, err := fa.getUserRecord(idToken)
	if err != nil {
		return "", core.NewAuthenticationFailedError(pkg, err)
	}

	return record.Email, nil
}

func (fa *FirebaseAuthenticator) RevokeRefreshTokens(idToken string) core.PortError {
	client, err := fa.app.Auth(context.Background())
	if err != nil {
		return core.NewConnectionError(pkg, err)
	}

	authToken, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		return core.NewAuthenticationFailedError(pkg, err)
	}

	err = client.RevokeRefreshTokens(context.Background(), authToken.UID)
	if err != nil {
		return core.NewAuthenticationFailedError(pkg, err)
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

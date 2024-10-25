package infrastructure

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

const (
	pkg string = "infrastructure.Firebase"
)

type (
	FirebaseAuthenticator struct {
		app *firebase.App
		log *zap.Logger
	}
)

func NewFirebaseAuthenticator(conf config.Firebase, log *zap.Logger) (*FirebaseAuthenticator, error) {
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
		log: log,
	}, nil

}

func (fa *FirebaseAuthenticator) CreateUser(id, email, password string) error {
	client, err := fa.app.Auth(context.Background())
	if err != nil {
		fa.logError("Firebase connection failed", err)
		return err
	}

	user := &auth.UserToCreate{}
	user.UID(id)
	user.Email(email)
	user.Password(password)

	_, err = client.CreateUser(context.Background(), user)
	if err != nil {
		fa.logError("Creating firebase user failed", err)
		return err
	}

	return nil
}

func (fa *FirebaseAuthenticator) GetUserID(idToken string) (string, error) {
	record, err := fa.getUserRecord(idToken)
	if err != nil {
		fa.logError("User retrieval failed", err)
		return "", err
	}

	return record.UID, nil
}

func (fa *FirebaseAuthenticator) GetUserEmail(idToken string) (string, error) {
	record, err := fa.getUserRecord(idToken)
	if err != nil {
		fa.logError("User retrieval failed", err)
		return "", err
	}

	return record.Email, nil
}

func (fa *FirebaseAuthenticator) RevokeRefreshTokens(idToken string) error {
	client, err := fa.app.Auth(context.Background())
	if err != nil {
		fa.logError("Firebase connection failed", err)
		return err
	}

	authToken, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		fa.logError("User verification failed", err)
		return port.ErrUserUnauthorized
	}

	return client.RevokeRefreshTokens(context.Background(), authToken.UID)
}

func (fa *FirebaseAuthenticator) getUserRecord(idToken string) (*auth.UserRecord, error) {
	client, err := fa.app.Auth(context.Background())
	if err != nil {
		fa.logError("Firebase connection failed", err)
		return nil, err
	}

	authToken, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		fa.logError("User verification failed", err)
		return nil, port.ErrUserUnauthorized
	}

	return client.GetUser(context.Background(), authToken.UID)
}

func (fa *FirebaseAuthenticator) logError(msg string, err error) {
	fa.log.Error(
		msg,
		zap.String("package", pkg),
		zap.String("reason", err.Error()),
	)
}

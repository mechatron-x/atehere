package notifier

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go/v4"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"google.golang.org/api/option"
)

type (
	FirestoreNotifier struct {
		app *firebase.App
	}
)

func NewFirestore(conf config.Firebase) (*FirestoreNotifier, error) {
	bytes, err := json.Marshal(conf)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsJSON(bytes)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return &FirestoreNotifier{
		app: app,
	}, nil
}

func (fen *FirestoreNotifier) NotifyOrderCreatedEvent(event *dto.OrderCreatedEvent) error {
	client, err := fen.app.Firestore(context.Background())
	if err != nil {
		return err
	}

	notificationData := map[string]interface{}{
		"invoke_time": event.InvokeTime,
		"message":     event.Message(),
		"table_name":  event.Table,
	}

	_, err = client.Collection(event.RestaurantID).
		Doc(event.ID.String()).
		Set(context.Background(), notificationData)

	return err
}

func (fen *FirestoreNotifier) NotifySessionClosedEvent(event *dto.SessionClosedEvent) error {
	client, err := fen.app.Firestore(context.Background())
	if err != nil {
		return err
	}

	notificationData := map[string]interface{}{
		"invoke_time": event.InvokeTime,
		"message":     event.Message(),
		"table_name":  event.Table,
	}

	_, err = client.Collection(event.RestaurantID).
		Doc(event.ID.String()).
		Set(context.Background(), notificationData)

	return err
}

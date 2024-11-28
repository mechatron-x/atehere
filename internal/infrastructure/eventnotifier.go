package infrastructure

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	FirebaseEventNotifier struct {
		app *firebase.App
	}
)

func NewFirecloudEventNotifier(conf config.Firebase) (*FirebaseEventNotifier, error) {
	bytes, err := json.Marshal(conf)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsJSON(bytes)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return &FirebaseEventNotifier{
		app: app,
	}, nil
}

func (fen *FirebaseEventNotifier) NotifyOrderCreatedEvent(event *dto.OrderCreatedEventView) error {
	client, err := fen.app.Firestore(context.Background())
	if err != nil {
		return err
	}

	notificationData := map[string]interface{}{
		"invoke_time": event.InvokeTime,
		"message":     event.Message(),
		"table_name":  event.Table,
	}

	_, err = client.Collection("session_events").
		Doc(event.RestaurantID).
		Update(context.Background(), []firestore.Update{
			{
				Path:  "events",
				Value: firestore.ArrayUnion(notificationData),
			},
		})
	if status.Code(err) == codes.NotFound {
		_, err := client.Collection("session_events").
			Doc(event.RestaurantID).
			Set(context.Background(), map[string]interface{}{
				"events": []interface{}{notificationData},
			})
		if err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

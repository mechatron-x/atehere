package ctx

import (
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/session/port"
	"github.com/mechatron-x/atehere/internal/session/service"
	"github.com/mechatron-x/atehere/internal/sqldb/repository"
	"gorm.io/gorm"
)

type Session struct {
	handler handler.Session
}

func NewSession(
	db *gorm.DB,
	authenticator port.Authenticator,
	eventNotifier port.EventNotifier,
) Session {
	repo := repository.NewSession(db)
	viewRepo := repository.NewSessionView(db)

	service := service.NewSession(
		repo,
		viewRepo,
		authenticator,
		eventNotifier,
		10,
	)

	handler := handler.NewSessionHandler(*service)
	return Session{
		handler: handler,
	}
}

func (s Session) Handler() handler.Session {
	return s.handler
}

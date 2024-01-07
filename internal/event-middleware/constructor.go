package eventmiddleware

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type EventService interface {
	RequestReceived(ctx context.Context, evt model.RequestReceived)
}

func New(eventService EventService) *Middleware {
	return &Middleware{
		EventService: eventService,
	}
}

type Middleware struct {
	EventService EventService
}

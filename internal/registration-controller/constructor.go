package registrationcontroller

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type MailService interface {
	Send(to string, subject string, contentType string, content string) error
}

type RegistrationService interface {
	Begin(ctx context.Context, ops model.BeginRegistrationOptions) (string, error)
	End(ctx context.Context, ops model.EndRegistrationOptions) (uint, error)
}

type SessionService interface {
	CreateWithAccountID(ctx context.Context, accountID uint) (uint, error)
	ReadByID(ctx context.Context, id uint) (model.Session, error)
}

func New(
	mailService MailService,
	registrationService RegistrationService,
	sessionService SessionService,
) *Controller {
	return &Controller{
		MailService:         mailService,
		RegistrationService: registrationService,
		SessionService:      sessionService,
	}
}

type Controller struct {
	MailService         MailService
	RegistrationService RegistrationService
	SessionService      SessionService
}

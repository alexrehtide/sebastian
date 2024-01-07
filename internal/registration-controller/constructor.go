package registrationcontroller

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type MailService interface {
	Send(to string, subject string, contentType string, content string) error
}

type RegistrationFormService interface {
	Begin(ctx context.Context, ops model.BeginRegistrationOptions) (string, error)
	End(ctx context.Context, ops model.EndRegistrationOptions) (uint, error)
}

type SessionService interface {
	CreateWithAccountID(ctx context.Context, accountID uint) (uint, error)
	ReadByID(ctx context.Context, id uint) (model.Session, error)
}

func New(
	mailService MailService,
	registrationFormService RegistrationFormService,
	sessionService SessionService,
) *Controller {
	return &Controller{
		MailService:             mailService,
		RegistrationFormService: registrationFormService,
		SessionService:          sessionService,
	}
}

type Controller struct {
	MailService             MailService
	RegistrationFormService RegistrationFormService
	SessionService          SessionService
}

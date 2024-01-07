package passwordresettingcontroller

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type MailService interface {
	Send(to string, subject string, contentType string, content string) error
}

type PasswordResettingService interface {
	Begin(ctx context.Context, ops model.BeginPasswordResettingOptions) (string, error)
	End(ctx context.Context, ops model.EndPasswordResettingOptions) error
}

func New(mailService MailService, passwordResettingService PasswordResettingService) *Controller {
	return &Controller{
		MailService:              mailService,
		PasswordResettingService: passwordResettingService,
	}
}

type Controller struct {
	MailService              MailService
	PasswordResettingService PasswordResettingService
}

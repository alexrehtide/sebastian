package authcontroller

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
)

type AccountProvider interface {
	Inject(ctx context.Context) *model.Account
}

type AuthService interface {
	Authenticate(ctx context.Context, in model.AuthenticateOptions) (model.Tokens, error)
}

type LoginAttemptService interface {
	CheckLoginAttempt(ctx context.Context, ip string) error
	FailLoginAttempt(ctx context.Context, ip string) error
	SuccessLoginAttempt(ctx context.Context, ip string) error
}

type MailService interface {
	Send(to string, subject string, contentType string, content string) error
}

type RBACService interface {
	ReadAccountRoles(ctx context.Context, accountID uint) ([]model.Role, error)
}

type RegistrationFormService interface {
	BeginRegistration(ctx context.Context, ops model.BeginRegistrationOptions) (string, error)
	EndRegistration(ctx context.Context, ops model.EndRegistrationOptions) (uint, error)
}

type SessionService interface {
	CreateWithAccountID(ctx context.Context, accountID uint) (uint, error)
	ReadByID(ctx context.Context, id uint) (model.Session, error)
}

func New(
	accountProvider AccountProvider,
	authService AuthService,
	loginAttemptService LoginAttemptService,
	mailService MailService,
	rbacService RBACService,
	registrationFormService RegistrationFormService,
	sessionService SessionService,
) *Controller {
	return &Controller{
		AccountProvider:         accountProvider,
		AuthService:             authService,
		LoginAttemptService:     loginAttemptService,
		MailService:             mailService,
		RBACService:             rbacService,
		RegistrationFormService: registrationFormService,
		SessionService:          sessionService,
	}
}

type Controller struct {
	AccountProvider         AccountProvider
	AuthService             AuthService
	LoginAttemptService     LoginAttemptService
	MailService             MailService
	RBACService             RBACService
	RegistrationFormService RegistrationFormService
	SessionService          SessionService
}

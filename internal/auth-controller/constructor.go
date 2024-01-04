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

type RBACService interface {
	ReadAccountRoles(ctx context.Context, accountID uint) ([]model.Role, error)
}

func New(
	accountProvider AccountProvider,
	authService AuthService,
	loginAttemptService LoginAttemptService,
	rbacService RBACService,
) *Controller {
	return &Controller{
		AccountProvider:     accountProvider,
		AuthService:         authService,
		LoginAttemptService: loginAttemptService,
		RBACService:         rbacService,
	}
}

type Controller struct {
	AccountProvider     AccountProvider
	AuthService         AuthService
	LoginAttemptService LoginAttemptService
	RBACService         RBACService
}

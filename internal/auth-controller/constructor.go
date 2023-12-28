package authcontroller

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type AccountProvider interface {
	Inject(c *fiber.Ctx) *model.Account
}

type AuthService interface {
	Authenticate(ctx context.Context, in model.AuthenticateOptions) (model.Tokens, error)
}

type RBACService interface {
	ReadAccountRoles(ctx context.Context, accountID uint) ([]model.Role, error)
}

func New(accountProvider AccountProvider, authService AuthService, rbacService RBACService) *Controller {
	return &Controller{
		AccountProvider: accountProvider,
		AuthService:     authService,
		RBACService:     rbacService,
	}
}

type Controller struct {
	AccountProvider AccountProvider
	AuthService     AuthService
	RBACService     RBACService
}

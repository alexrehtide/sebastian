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

func New(accountProvider AccountProvider, authService AuthService) *Controller {
	return &Controller{
		AccountProvider: accountProvider,
		AuthService:     authService,
	}
}

type Controller struct {
	AccountProvider AccountProvider
	AuthService     AuthService
}

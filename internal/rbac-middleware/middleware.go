package rbacmiddleware

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type AccountProvider interface {
	Inject(c *fiber.Ctx) *model.Account
}

type RBACService interface {
	AccountHasPermission(ctx context.Context, accountID uint, permission model.Permission) (bool, error)
}

func New(accountProvider AccountProvider, rbacService RBACService) *Middleware {
	return &Middleware{
		AccountProvider: accountProvider,
		RBACService:     rbacService,
	}
}

type Middleware struct {
	AccountProvider AccountProvider
	RBACService     RBACService
}

func (m *Middleware) WithPermission(permission model.Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

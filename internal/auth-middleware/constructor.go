package authmiddleware

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type AccountProvider interface {
	Provide(c *fiber.Ctx, acc model.Account)
}

func New(accountProvider AccountProvider) *Middleware {
	return &Middleware{AccountProvider: accountProvider}
}

type Middleware struct {
	AccountProvider AccountProvider
}

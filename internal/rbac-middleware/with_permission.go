package rbacmiddleware

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) WithPermission(permission model.Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

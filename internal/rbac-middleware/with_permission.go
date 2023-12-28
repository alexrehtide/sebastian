package rbacmiddleware

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) WithPermission(permission model.Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		account := m.AccountProvider.Inject(c)
		if account == nil {
			return c.Status(fiber.StatusForbidden).SendString("permission denied")
		}
		hasAccess, err := m.RBACService.AccountHasPermission(c.Context(), account.ID, permission)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		if !hasAccess {
			return c.Status(fiber.StatusForbidden).SendString("permission denied")
		}
		return c.Next()
	}
}

package authmiddleware

import (
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Authorize(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Next()
	}
	session, err := m.SessionService.ReadByAccessToken(c.Context(), token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("session not found")
	}
	if err := m.SessionService.Verify(session); err != nil {
		return c.Next()
	}
	m.SessionProvider.Provide(c, session)
	account, err := m.AccountService.ReadByID(c.Context(), session.AccountID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("account not found")
	}
	m.AccountProvider.Provide(c, account)
	return c.Next()
}

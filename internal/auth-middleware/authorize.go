package authmiddleware

import (
	middlewareerror "github.com/alexrehtide/sebastian/internal/middleware-error"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Authorize(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Next()
	}
	session, err := m.SessionService.ReadByAccessToken(c.UserContext(), token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(middlewareerror.ErrSessionNotFound.Error())
	}
	if err := m.SessionService.Verify(session); err != nil {
		return c.Next()
	}
	c.SetUserContext(m.SessionProvider.Provide(c.UserContext(), session))
	account, err := m.AccountService.ReadByID(c.UserContext(), session.AccountID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(middlewareerror.ErrAccountNotFound.Error())
	}
	c.SetUserContext(m.AccountProvider.Provide(c.UserContext(), account))
	return c.Next()
}

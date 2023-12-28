package authmiddleware

import "github.com/gofiber/fiber/v2"

func (m *Middleware) Authorize(c *fiber.Ctx) error {
	return c.Next()
}

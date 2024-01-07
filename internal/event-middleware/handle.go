package eventmiddleware

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Handle(c *fiber.Ctx) error {
	e := model.RequestReceived{
		IP:   c.IP(),
		Body: string(c.Body()),
		Path: c.Path(),
	}

	m.EventService.RequestReceived(c.UserContext(), e)

	return c.Next()
}

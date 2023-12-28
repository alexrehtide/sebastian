package sessionprovider

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

func (p *Provider) Inject(c *fiber.Ctx) *model.Session {
	acc, ok := c.Locals(SESSION_INJECT_KEY).(model.Session)
	if !ok {
		return nil
	}
	return &acc
}

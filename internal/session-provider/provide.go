package sessionprovider

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

func (p *Provider) Provide(c *fiber.Ctx, acc model.Session) {
	c.Locals(SESSION_INJECT_KEY, acc)
}

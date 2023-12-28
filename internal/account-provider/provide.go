package accountprovider

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

func (p *Provider) Provide(c *fiber.Ctx, acc model.Account) {
	c.Locals(ACCOUNT_INJECT_KEY, acc)
}

package accountprovider

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

func (p *Provider) Inject(c *fiber.Ctx) *model.Account {
	acc, ok := c.Locals(ACCOUNT_INJECT_KEY).(model.Account)
	if !ok {
		return nil
	}
	return &acc
}

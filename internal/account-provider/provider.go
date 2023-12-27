package accountprovider

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

const (
	ACCOUNT_INJECT_KEY = "account"
)

func New() *Provider {
	return &Provider{}
}

type Provider struct{}

func (p *Provider) Provide(c *fiber.Ctx, acc model.Account) {
	c.Locals(ACCOUNT_INJECT_KEY, acc)
}

func (p *Provider) Inject(c *fiber.Ctx) *model.Account {
	acc, ok := c.Locals(ACCOUNT_INJECT_KEY).(model.Account)
	if !ok {
		return nil
	}
	return &acc
}

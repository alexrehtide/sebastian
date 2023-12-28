package authcontroller

import "github.com/gofiber/fiber/v2"

type AuthorizeOutput struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func (ctrl *Controller) Authorize(c *fiber.Ctx) error {
	account := ctrl.AccountProvider.Inject(c)
	return c.JSON(AuthorizeOutput{
		ID:    account.ID,
		Email: account.Email,
	})
}

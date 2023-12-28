package authcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type AuthorizeOutput struct {
	ID    uint         `json:"id"`
	Email string       `json:"email"`
	Roles []model.Role `json:"roles"`
}

func (ctrl *Controller) Authorize(c *fiber.Ctx) error {
	account := ctrl.AccountProvider.Inject(c)
	roles, err := ctrl.RBACService.ReadAccountRoles(c.Context(), account.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(AuthorizeOutput{
		ID:    account.ID,
		Email: account.Email,
		Roles: roles,
	})
}

package accountcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type UpdateInput struct {
	ID    uint   `params:"id"`
	Email string `json:"email"`
}

func (ctrl *Controller) Update(c *fiber.Ctx) error {
	var input UpdateInput
	if err := c.ParamsParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err := ctrl.AccountService.Update(
		c.Context(),
		input.ID,
		model.UpdateAccountOptions{
			Email: input.Email,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

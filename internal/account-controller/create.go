package accountcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type CreateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateOutput struct {
	uint
}

func (ctrl *Controller) Create(c *fiber.Ctx) error {
	var input CreateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	output, err := ctrl.AccountService.Create(
		c.Context(),
		model.CreateAccountOptions{
			Email:    input.Email,
			Password: input.Password,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(CreateOutput{output})
}

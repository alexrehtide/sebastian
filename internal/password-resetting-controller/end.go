package passwordresettingcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/password"
	"github.com/gofiber/fiber/v2"
)

type EndInput struct {
	ResettingCode string `json:"resettingCode"`
	NewPassword   string `json:"newPassword"`
}

func (ctrl *Controller) End(c *fiber.Ctx) error {
	var input EndInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err := ctrl.PasswordResettingService.End(c.UserContext(), model.EndPasswordResettingOptions{
		ResettingCode: input.ResettingCode,
		NewPassword:   password.HashPassword(input.NewPassword),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

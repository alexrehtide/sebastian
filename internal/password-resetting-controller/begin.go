package passwordresettingcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type BeginInput struct {
	Email string `json:"email"`
}

func (ctrl *Controller) Begin(c *fiber.Ctx) error {
	var input BeginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	url, err := ctrl.PasswordResettingService.Begin(c.UserContext(), model.BeginPasswordResettingOptions{Email: input.Email})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err = ctrl.MailService.Send(input.Email, "Password resetting", "text/plain", url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

package registrationcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/alexrehtide/sebastian/pkg/password"
	"github.com/gofiber/fiber/v2"
)

type BeginRegistrationInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type BeginRegistrationOutput struct {
	CreatedID uint `json:"createdId"`
}

func (ctrl *Controller) Begin(c *fiber.Ctx) error {
	var input BeginRegistrationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	url, err := ctrl.RegistrationService.Begin(c.UserContext(), model.BeginRegistrationOptions{
		Email:    input.Email,
		Username: input.Username,
		Password: password.HashPassword(input.Password),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err = ctrl.MailService.Send(input.Email, "Registration", "text/plain", url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

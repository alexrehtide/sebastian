package authcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type EndRegistrationInput struct {
	VerificationCode string `json:"verificationCode"`
}

type EndRegistrationOutput struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (ctrl *Controller) EndRegistration(c *fiber.Ctx) error {
	var input EndRegistrationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	accID, err := ctrl.RegistrationFormService.EndRegistration(c.UserContext(), model.EndRegistrationOptions{VerificationCode: input.VerificationCode})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	sessionID, err := ctrl.SessionService.CreateWithAccountID(c.UserContext(), accID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	session, err := ctrl.SessionService.ReadByID(c.UserContext(), sessionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(EndRegistrationOutput{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
	})
}

package authcontroller

import (
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type AuthenticateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateOutput struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (ctrl *Controller) Authenticate(c *fiber.Ctx) error {
	ip := c.IP()
	if err := ctrl.AuthService.CheckLoginAttempt(c.UserContext(), ip); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
	var input AuthenticateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	tokens, err := ctrl.AuthService.Authenticate(
		c.UserContext(),
		model.AuthenticateOptions{
			Email:    input.Email,
			Password: input.Password,
		},
	)
	if err != nil {
		if err := ctrl.AuthService.FailLoginAttempt(c.UserContext(), ip); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := ctrl.AuthService.SuccessLoginAttempt(c.UserContext(), ip); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(AuthenticateOutput{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

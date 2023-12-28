package authcontroller

import (
	"context"

	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
)

type AccountProvider interface {
	Inject(c *fiber.Ctx) *model.Account
}

type AuthService interface {
	Authenticate(ctx context.Context, in model.AuthenticateOptions) (model.Tokens, error)
}

func New(accountProvider AccountProvider, authService AuthService) *Controller {
	return &Controller{
		AccountProvider: accountProvider,
		AuthService:     authService,
	}
}

type Controller struct {
	AccountProvider AccountProvider
	AuthService     AuthService
}

type AuthenticateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateOutput struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (ctrl *Controller) Authenticate(c *fiber.Ctx) error {
	var input AuthenticateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	tokens, err := ctrl.AuthService.Authenticate(
		c.Context(),
		model.AuthenticateOptions{
			Email:    input.Email,
			Password: input.Password,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(AuthenticateOutput{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

func (ctrl *Controller) Authorize(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Unimplemented")
}

func (ctrl *Controller) Logout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Unimplemented")
}

func (ctrl *Controller) Refresh(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).SendString("Unimplemented")
}

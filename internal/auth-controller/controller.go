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
	Authenticate(ctx context.Context, in model.AuthenticateInput) (model.AuthenticateOutput, error)
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

func (ctrl *Controller) Authenticate(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	output, err := ctrl.AuthService.Authenticate(
		c.Context(),
		model.AuthenticateInput{
			Email:    input.Email,
			Password: input.Password,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(output)
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
